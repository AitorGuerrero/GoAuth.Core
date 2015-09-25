package user

import (
	"code.google.com/p/go-uuid/uuid"
)
func GenerateNewToken (user User) Token {
	return Token{Code(uuid.NewRandom()), user}
}

type TokenSource interface {
	Add(t Token) error
	ByUser (u User) (Token, error)
}

type Code uuid.UUID

type Token struct {
	code Code
	user User
}

func (t Token) Code() Code {
	return t.code
}

func (t Token) User() User {
	return t.user
}
