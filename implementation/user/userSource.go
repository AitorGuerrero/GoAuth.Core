package user

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Source struct {
	Collection map[string]*user.User
}

func (s *Source) Add (u *user.User) error {
	if s.Collection[string(u.Id())] != nil {
		return user.DuplicatedIdError{u.Id()}
	}
	s.Collection[string(u.Id())] = u

	return nil
}

func (us Source) Get (i user.Id) (u *user.User, err error) {
	u = us.Collection[string(i)]
	if u == nil {
		err = user.NotExistentUser{i}
	}

	return u, err
}
