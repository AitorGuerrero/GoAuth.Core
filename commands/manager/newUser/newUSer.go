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
		return ManagerDoesNotExistError{err}
	}
	if uid, err = user.ParseId(r.UserId); nil != err {
		return MalformedUserIdError{err}
	}
	u := c.Factory.Make(uid, r.Passkey)
	err = c.UserSource.Add(&u)
	if _, ok := err.(user.DuplicatedIdError); ok  {
		return DuplicatedIdError{err}
	}
	m.AddUser(&u)

	return nil;
}
