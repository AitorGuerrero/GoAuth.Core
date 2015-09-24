package newUser

import (
	"errors"
	"github.com/AitorGuerrero/UserGo/user"
	//	"fmt"
)

type Command struct {
	source user.UserSource
	encryptor user.PasskeyEncryptor
}

type Request struct {
	Id string
	passkey string
}

func (c Command) Execute(r Request) (error) {
	if "" == r.Id {
		return errors.New("user should have an identifier")
	}
	var userId = user.Id(r.Id);
	u := user.New(userId, c.encryptor.Encrypt(userId, user.Passkey(r.passkey)))
	c.source.Add(u)

	return nil;
}
