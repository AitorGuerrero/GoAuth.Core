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
	var userId = user.Id(r.Id);
	u := user.New(userId, c.Encryptor.Encrypt(userId, user.Passkey(r.Passkey)))
	c.Source.Add(u)

	return nil;
}
