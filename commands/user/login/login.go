package login

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Login user.Login
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
	id := user.Id(req.Id)
	p := user.Passkey(req.Passkey)
	t, err := c.Login.Try(id, p)
	res.SessionToken = string(t.Code)

	return res, err
}
