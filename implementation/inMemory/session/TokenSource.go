package session

import (
	"github.com/AitorGuerrero/UserGo/session"
	"github.com/AitorGuerrero/UserGo/user"

	"errors"
)

type TokenSource struct {
	collection []session.Token
}

func (s *TokenSource) Add(t session.Token) error {
	s.collection = append(s.collection, t)

	return nil
}

func (s TokenSource) ByUser (u user.User) (session.Token, error) {
	for _, t := range s.collection {
		if t.User() == u {
			return t, nil
		}
	}

	return session.Token{}, errors.New("Not found a token for the user")
}
