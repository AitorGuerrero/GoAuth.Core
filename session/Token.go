package session

import (
	"github.com/AitorGuerrero/UserGo/user"
)
func GenerateNewToken (user user.User) Token {
	return Token{"newToken", user}
}

type TokenSource interface {
	Add(t Token) error
	ByUser (u user.User) (Token, error)
}

type Token struct {
	code string
	user user.User
}

func (t Token) Code() string {
	return t.code
}

func (t Token) User() user.User {
	return t.user
}
