package login

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	validator user.SignInValidator
}

type Request struct {
	Id string
	passkey string
}

type Response struct {
	SessionToken string
}

func (c Command) Execute(r Request) (Response, error) {
	res := Response{}
	error := c.validator.Validate(user.Id(r.Id), user.Passkey(r.passkey))
	if nil != error {
		return res, error
	}
	res.SessionToken = "newToken"

	return res, nil;
}
