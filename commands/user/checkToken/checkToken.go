package checkToken

import (
	"github.com/AitorGuerrero/UserGo"
	"github.com/AitorGuerrero/UserGo/user"
	"errors"
)
type Command struct {
	TokenSource UserGo.TokenSource
	UserSource user.UserSource
}

func (c Command) Execute(req Request) (error) {
	user, _ := c.UserSource.ById(user.Id(req.UserId))
	token, err := c.TokenSource.ByUser(user)
	if (nil != err || token.User() != user) {
		return errors.New("Incorrect Token")
	}

	return nil
}

type Request struct {
	UserId string
	Token string
}
