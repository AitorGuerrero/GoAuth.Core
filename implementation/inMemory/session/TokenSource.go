package session

import (
	"github.com/AitorGuerrero/UserGo/user"

	"errors"
)

type TokenSource struct {
	collection []user.Token
}

func (s *TokenSource) Add(t user.Token) error {
	s.collection = append(s.collection, t)

	return nil
}

func (s TokenSource) ByUser (u user.User) (user.Token, error) {
	for _, t := range s.collection {
		if t.User() == u {
			return t, nil
		}
	}

	return user.Token{}, errors.New("Not found a token for the user")
}
