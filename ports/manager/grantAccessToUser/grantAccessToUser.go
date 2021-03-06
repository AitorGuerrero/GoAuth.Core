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
	var mi, ui user.Id
	mi = user.ParseId(r.ManagerId);
	ui = user.ParseId(r.UserId);
	m, err := c.ManagerSource.Get(mi)
	if _, ok := err.(manager.NotExistsError); ok {
		return ManagerDoesNotExist{err, req.ManagerId}
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
