package session

import (
	"github.com/AitorGuerrero/UserGo/user"
	"code.google.com/p/go-uuid/uuid"
)
func GenerateNewToken (user user.User) Token {
	return Token{Code(uuid.NewRandom()), user}
}

type TokenSource interface {
	Add(t Token) error
	ByUser (u user.User) (Token, error)
}

type Code uuid.UUID

type Token struct {
	code Code
	user user.User
}

func (t Token) Code() Code {
	return t.code
}

func (t Token) User() user.User {
	return t.user
}
