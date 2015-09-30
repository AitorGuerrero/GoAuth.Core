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
	var m *manager.Manager
	var uid user.Id
	mi, _ := user.ParseId(r.ManagerId)
	if m, err = c.ManagerSource.Get(mi); nil != err {
		return ManagerDoesNotExistError{}
	}
	if uid, err = user.ParseId(r.UserId); nil != err {
		return MalformedUserIdError{}
	}
	u := c.Factory.Make(uid, r.Passkey)
	c.UserSource.Add(&u)
	m.AddUser(&u)

	return nil;
}
