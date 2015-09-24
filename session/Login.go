package session

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Login struct {
	UserSource user.UserSource
	Validator user.SignInValidator
	TokensSource TokenSource
}

func (l Login) Try(id user.Id, p user.Passkey) (Token, error) {
	err := l.Validator.Validate(id, p)
	if nil != err {
		return Token{}, err
	}
	u, _ := l.UserSource.ById(id)
	t, err := l.TokensSource.ByUser(u)
	if (nil != err) {
		t = GenerateNewToken(u)
		l.TokensSource.Add(t)
	}

	return t, nil
}
