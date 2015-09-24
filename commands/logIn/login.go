package login

import (
	"github.com/AitorGuerrero/UserGo/user"
	"errors"
//	"fmt"
)


type Command struct {
	source user.UserSource
	encryptor user.PasskeyEncryptor
}

type Request struct {
	Id user.Id
	passkey user.Passkey
}

func (c Command) Execute(r Request) (error) {
	u, err := c.source.ById(r.Id)
	if (nil != err) {
		return err
	}
	if c.encryptor.Encrypt(r.Id, r.passkey) != u.EncryptedPasskey() {
		return errors.New("The password doen't match")
	}
	return nil;
}
