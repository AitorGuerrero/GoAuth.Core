package checkToken

import (
	"github.com/AitorGuerrero/UserGo/user"
)
type Command struct {
	TokenChecker user.TokenChecker
	UserSource user.UserSource
}

func (c Command) Execute(req Request) (error) {
	u, err := c.UserSource.Get(user.Id(req.UserId))
	if (nil == err) {
		t := user.Token{user.TokenCode(req.Token), u}
		err = c.TokenChecker.Check(t)
	}

	return err
}

type Request struct {
	UserId string
	Token string
}
