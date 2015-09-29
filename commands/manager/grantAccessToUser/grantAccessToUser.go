package grantAccessToUser

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)

type Command struct {
	UserSource user.Source
	ManagerSource manager.ManagerSource
}

type Request struct {
	ManagerId string
	Namespace string
}

func (c Command) Execute(r Request) (err error) {
	m, err := c.ManagerSource.ById(r.ManagerId)
	if nil != err {
		return
	}
	m.GrantAccessToUser(u, user.Namespace(r.Namespace))

	return
}
