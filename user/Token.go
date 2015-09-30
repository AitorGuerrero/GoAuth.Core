package user

import (
	"code.google.com/p/go-uuid/uuid"
)

type TokenCode uuid.UUID

type Token struct {
	Code TokenCode
}

func (u User) Token() Token {
	return u.token
}

func (u User) HasToken() bool {
	return "" != string(u.token.Code)
}

func (u *User) GenerateToken() {
	u.token = Token{TokenCode(uuid.NewRandom())}
}

func (t Token) IsSame(t2 Token) bool {
	return string(t.Code) == string(t2.Code)
}
