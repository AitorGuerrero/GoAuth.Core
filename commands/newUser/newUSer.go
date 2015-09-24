package newUser

import (
	"errors"
	"github.com/AitorGuerrero/UserGo/user"
	//	"fmt"
)

type Command struct {
	source user.UserSource
}

type Request struct {
	Id user.Id
}

func (c Command) Execute(r Request) (error) {
	if "" == r.Id {
		return errors.New("user should have an identifier")
	}

	u := user.New(r.Id)
	c.source.Add(u)

	return nil;
}
