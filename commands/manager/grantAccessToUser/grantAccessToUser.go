package grantAccessToUser

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)

type Command struct {
	UserSource user.Source
	ManagerSource manager.Source
}

type Request struct {
	ManagerId string
	UserId string
	Namespace string
}

func (c Command) Execute(r Request) (err error) {
	mi, err := user.ParseId(r.ManagerId)
	if (nil != err) {
		return
	}
	m, err := c.ManagerSource.Get(mi)
	if _, ok := err.(manager.NotExistsError); ok {
		return ManagerDoesNotExist{err, req.ManagerId}
	}
	ui, err := user.ParseId(r.UserId)
	if (nil != err) {
		return
	}
	u, err := c.UserSource.Get(ui)
	if _, ok := err.(user.NotExistentUser); ok {
		return UserDoesNotExist{err, string(ui)}
	}
	err = m.GrantAccessToUser(u, user.Namespace(r.Namespace))
	if _, ok := err.(manager.DoesNotOwnTheUser); ok {
		return ManagerDoesNotOwnTheUser{err, req.UserId, req.ManagerId}
	}
	if _, ok := err.(manager.ManagerDoesNotOwnTheNamespace); ok {
		return ManagerDoesNotOwnTheNamespace{err, req.ManagerId, req.Namespace}
	}

	return
}
