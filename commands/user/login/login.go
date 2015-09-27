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
	u2, err2 := c.Login.Try(u, user.Passkey(req.Passkey))
	c.UserSource.Persist(u2)
	res.SessionToken = string(u2.Token().Code)

	return res, err2
}
