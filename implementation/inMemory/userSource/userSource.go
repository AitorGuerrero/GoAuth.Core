package userSource

import (
	"github.com/AitorGuerrero/UserGo/user"
	"errors"
//	"fmt"
)

func New() userSource {
	return userSource{};
}

type userSource struct {
	collection []user.User
}

func (us *userSource) Add (u user.User) error {
	us.collection = append(us.collection, u)

	return nil
}

func (us userSource) ById (i user.Id) (user.User, error) {
	for _, u := range us.collection {
		if u.Id() == i {
			return u, nil
		}
	}

	return user.User{}, errors.New("Not found user "+string(i))
}
