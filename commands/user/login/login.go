package login

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Login user.Login
	UserSource user.UserSource
}

type Request struct {
	Id string
	Passkey string
}

type Response struct {
	SessionToken string
}

func (c Command) Execute(req Request) (Response, error) {
	res := Response{}
	u, err := c.UserSource.Get(user.Id(req.Id))
	if nil != err {
		return res, err
	}
	t, err := c.Login.Try(u, user.Passkey(req.Passkey))
	res.SessionToken = string(t.Code)

	return res, err
}
