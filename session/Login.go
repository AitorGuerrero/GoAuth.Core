package session

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo"
)

type Login struct {
	UserSource user.UserSource
	Validator user.SignInValidator
	TokensSource UserGo.TokenSource
}

func (l Login) Try(id user.Id, p user.Passkey) (UserGo.Token, error) {
	t := UserGo.Token{}
	u, err := l.UserSource.ById(id)
	if nil != err {
		return t, err
	}
	err = l.Validator.Validate(u, p)
	if nil != err {
		return t, err
	}
	t, err = l.TokensSource.ByUser(u)
	if (nil != err) {
		t = UserGo.GenerateNewToken(u)
		l.TokensSource.Add(t)
	}

	return t, nil
}
