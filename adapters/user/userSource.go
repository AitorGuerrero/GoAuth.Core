package user

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Source struct {
	Collection map[string]*user.User
}

func NewSource() Source {
	return Source{map[string]*user.User{}}
}

func (s *Source) Add (u *user.User) error {
	if s.Collection[u.Id().Serialize()] != nil {
		return user.DuplicatedIdError{u.Id()}
	}
	s.Collection[u.Id().Serialize()] = u

	return nil
}

func (us Source) Get (i user.Id) (u *user.User, err error) {
	u = us.Collection[i.Serialize()]
	if u == nil {
		err = user.NotExistentUser{i}
	}

	return u, err
}
