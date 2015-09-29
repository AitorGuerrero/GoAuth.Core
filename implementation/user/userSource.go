package user

import (
	"github.com/AitorGuerrero/UserGo/user"
	"errors"
)

type Source struct {
	Collection map[string]user.User
}

func (us *Source) Add (u user.User) error {
	if us.Collection[string(u.Id())].Id().Equal(u.Id()) {
		return errors.New("Existing user")
	}
	us.Persist(u)

	return nil
}

func (us Source) Get (i user.Id) (user.User, error) {
	u := us.Collection[string(i)]
	if !u.Id().Equal(i) {
		return u, errors.New("Invalid User id")
	}

	return u, nil
}

func (us *Source) Persist (u user.User) error {
	us.Collection[string(u.Id())] = u

	return nil
}
