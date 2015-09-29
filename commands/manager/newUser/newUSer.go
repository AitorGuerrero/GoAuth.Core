package newUser

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)

type Command struct {
	ManagerSource manager.Source
	UserSource user.Source
	Factory user.Factory
}

type Request struct {
	ManagerId string
	UserId string
	Passkey string
}

func (c Command) Execute(r Request) (err error) {
	m, err := c.ManagerSource.Get(user.Id(r.ManagerId))
	if nil != err {
		return
	}
	u := c.Factory.Make(user.Id(r.UserId), user.Passkey(r.Passkey))
	c.UserSource.Add(u)
	m.AddUser(u)
	c.ManagerSource.Persist(m)

	return nil;
}
