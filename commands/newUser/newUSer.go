package newUser

import (
	"errors"
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Source user.UserSource
	Factory user.Factory
}

type Request struct {
	Id string
	Passkey string
}

func (c Command) Execute(r Request) (error) {
	if "" == r.Id {
		return errors.New("user should have an identifier")
	}
	uid := user.Id(r.Id);
	p := user.Passkey(r.Passkey)
	u := c.Factory.Make(uid, p)
	c.Source.Add(u)

	return nil;
}
