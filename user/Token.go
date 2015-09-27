package user

import (
	"code.google.com/p/go-uuid/uuid"
)
func GenerateNewToken (user User) Token {
	return Token{TokenCode(uuid.NewRandom()), user}
}

type TokenSource interface {
	Add(t Token) error
	ByUser (u User) (Token, error)
}

type TokenCode uuid.UUID

type Token struct {
	Code TokenCode
	User User
}
