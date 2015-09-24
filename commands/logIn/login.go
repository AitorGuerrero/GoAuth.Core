package login

import (
	"github.com/AitorGuerrero/UserGo/user"
	"errors"
//	"fmt"
)

type Request struct {
	id user.Id
	passkey user.Passkey
}

type Command struct {
	source user.UserSource
}

func (c Command) Execute(r Request) (error) {
	u, err := c.source.ById(r.id)
	if (nil != err) {
		return err
	}
	if r.passkey != u.Passkey() {
		return errors.New("The password doen't match")
	}
	return nil;
}
