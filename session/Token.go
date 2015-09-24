package session

import (
	"github.com/AitorGuerrero/UserGo/user"

	"errors"
)
func GenerateNewToken (user user.User) Token {
	return Token{"newToken", user}
}

type TokenSource struct {
	collection []Token
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

func (s *TokenSource) Add(t Token) error {
	s.collection = append(s.collection, t)

	return nil
}

func (s TokenSource) ByUser (u user.User) (Token, error) {
	for _, t := range s.collection {
		if t.user == u {
			return t, nil
		}
	}

	return Token{}, errors.New("Not found a token for the user")
}
