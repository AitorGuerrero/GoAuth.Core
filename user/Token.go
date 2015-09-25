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
	Code Code
	User User
}
