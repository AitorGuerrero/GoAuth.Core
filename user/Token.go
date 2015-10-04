package user

import (
	"code.google.com/p/go-uuid/uuid"
)

type Token uuid.UUID

func (t Token) IsSame(t2 Token) bool {
	return uuid.Equal(uuid.UUID(t), uuid.UUID(t2))
}

func (t Token) IsEmpty() bool {
	return t.Serialize() == ""
}

func (t Token) Serialize() string {
	return uuid.UUID(t).String()
}

func ParseToken(s string) Token {
	return Token(uuid.Parse(s))
}

func GenerateToken() Token {
	return Token(uuid.NewRandom())
}

