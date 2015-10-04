package login

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Login user.Login
	UserSource user.Source
}

type Request struct {
	Id string
	Passkey string
	Namespace string
}

type Response struct {
	SessionToken string
}

func (c Command) Execute(req Request) (res Response, err error) {
	u, err := c.Login.Try(
		user.ParseId(req.Id),
		req.Passkey,
		user.Namespace(req.Namespace),
	)

	if _, ok := err.(user.NotExistentUser); ok {
		return res, UserDoesNotExist{err, req.Id}
	}
	if _, ok := err.(user.IncorrectPasskeyError); ok {
		return res, IncorrectPasskeyError{err, req.Passkey}
	}
	if _, ok := err.(user.IncorrectNamespaceError); ok {
		return res, IncorrectNamespaceError{err}
	}
	if(nil != err) {
		return
	}
	res.SessionToken = u.Token.Serialize()

	return
}
