package login

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/session"
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

type Response struct {
	SessionToken session.Token
}

func (c Command) Execute(r Request) (Response, error) {
	res := Response{}
	u, err := c.source.ById(r.Id)
	if (nil != err) {
		return res, err
	}
	if c.encryptor.Encrypt(r.Id, r.passkey) != u.EncryptedPasskey() {
		return res, errors.New("The password doen't match")
	}

	res.SessionToken = session.Token("newToken")

	return res, nil;
}
