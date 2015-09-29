package newUser

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)

type Command struct {
	ManagerSource manager.Source
	UserSource user.UserSource
	Factory user.Factory
}

type Request struct {
	ManagerId string
	UserId string
	Passkey string
}

func (c Command) Execute(r Request) (err error) {
	m, err := c.ManagerSource.ById(r.ManagerId)
	if nil != err {
		return
	}
	uid := user.Id(r.Id);
	p := user.Passkey(r.Passkey)
	u := c.Factory.Make(uid, p)
	c.UserSource.Add(u)
	m.AddUser(u)
	c.ManagerSource.Persist(m)

	return nil;
}
