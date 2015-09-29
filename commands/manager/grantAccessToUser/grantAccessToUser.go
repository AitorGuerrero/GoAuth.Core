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
	m, err := c.ManagerSource.ById(user.Id(r.ManagerId))
	if nil != err {
		return
	}
	u, err := c.UserSource.Get(user.Id(r.UserId))
	if nil != err {
		return
	}
	m.GrantAccessToUser(u, user.Namespace(r.Namespace))

	return
}
