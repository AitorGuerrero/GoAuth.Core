package checkToken

import (
	"github.com/AitorGuerrero/UserGo/user"
)
type Command struct {
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

	if pid = user.ParseId(req.UserId); nil == pid {
		return InvalidIdError{}
	}
	if u, err = c.UserSource.Get(pid); nil != err {
		return InvalidIdError{}
	}
	n := user.Namespace(req.Namespace)
	if !u.CheckToken(user.ParseToken(req.Token), n) {
		return IncorrectTokenError{}
	}

	return err
}
