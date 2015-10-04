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
	var pid user.Id
	var u *user.User
	var err error

	if pid, err = user.ParseId(req.UserId); nil != err {
		return InvalidIdError{}
	}
	if u, err = c.UserSource.Get(pid); nil != err {
		return InvalidIdError{}
	}
	n := user.Namespace(req.Namespace)
	if err = c.TokenChecker.Check(*u, user.ParseToken(req.Token), n,); nil != err {
		if _, ok := err.(user.IncorrectTokenError); ok {
			return IncorrectTokenError{}
		}
		if _, ok := err.(user.AccessErrorToNamespace); ok {
			return AccessErrorToNamespace{}
		}
	}

	return err
}
