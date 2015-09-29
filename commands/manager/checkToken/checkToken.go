package checkToken

import (
	"github.com/AitorGuerrero/UserGo/user"
)
type Command struct {
	TokenChecker user.TokenChecker
	UserSource user.Source
}

type Request struct {
	UserId string
	Token string
	Namespace string
}

func (c Command) Execute(req Request) (error) {
	u, err := c.UserSource.Get(user.Id(req.UserId))
	if (nil != err) {
		return err
	}
	t := user.Token{user.TokenCode(req.Token)}
	n := user.Namespace(req.Namespace)
	return c.TokenChecker.Check(*u, t, n)
}
