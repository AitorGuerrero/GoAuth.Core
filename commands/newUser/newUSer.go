package newUser

import (
	"errors"
	"github.com/AitorGuerrero/UserGo/user"
	//	"fmt"
)

type Command struct {
	Source user.UserSource
	Encryptor user.PasskeyEncryptor
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
	ep := c.Encryptor.Encrypt(uid, p)
	u := user.New(uid, ep)
	c.Source.Add(u)

	return nil;
}
