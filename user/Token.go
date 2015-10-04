package user

import (
	"code.google.com/p/go-uuid/uuid"
)

type TokenCode uuid.UUID

type Token struct {
	Code TokenCode
}

func (t Token) IsSame(t2 Token) bool {
	return string(t.Code) == string(t2.Code)
}

func (t Token) IsEmpty() bool {
	return string(t.Code) == ""
}

func GenerateToken() Token {
	return Token{TokenCode(uuid.NewRandom())}
}
