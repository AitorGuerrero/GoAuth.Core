package user

import (
	"github.com/AitorGuerrero/UserGo/user"
	"errors"
)

type UserSource struct {
	Collection map[user.Id]user.User
}

func (us *UserSource) Add (u user.User) error {
	if us.Collection[u.Id()].Id() == u.Id() {
		return errors.New("Existing user")
	}
	us.Persist(u)

	return nil
}

func (us *UserSource) Get (i user.Id) (user.User, error) {
	u := us.Collection[i]
	if u.Id() != i {
		return u, errors.New("Invalid User id")
	}

	return u, nil
}

func (us *UserSource) Persist (u user.User) error {
	us.Collection[u.Id()] = u

	return nil
}
