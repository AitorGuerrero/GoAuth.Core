package login

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/session"
)

type Command struct {
	UserSource user.UserSource
	Validator user.SignInValidator
	TokensSource session.TokenSource
}

type Request struct {
	Id string
	Passkey string
}

type Response struct {
	SessionToken string
}

func (c Command) Execute(r Request) (Response, error) {
	res := Response{}
	error := c.Validator.Validate(user.Id(r.Id), user.Passkey(r.Passkey))
	if nil != error {
		return res, error
	}
	u, _ := c.UserSource.ById(user.Id(r.Id))
	token := session.GenerateNewToken(u)
	c.TokensSource.Add(token)
	res.SessionToken = token.Code()

	return res, nil;
}
