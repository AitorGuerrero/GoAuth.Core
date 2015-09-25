package session

import (
	"github.com/AitorGuerrero/UserGo"
	"github.com/AitorGuerrero/UserGo/user"

	"errors"
)

type TokenSource struct {
	collection []UserGo.Token
}

func (s *TokenSource) Add(t UserGo.Token) error {
	s.collection = append(s.collection, t)

	return nil
}

func (s TokenSource) ByUser (u user.User) (UserGo.Token, error) {
	for _, t := range s.collection {
		if t.User() == u {
			return t, nil
		}
	}

	return UserGo.Token{}, errors.New("Not found a token for the user")
}
