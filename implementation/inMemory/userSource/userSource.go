package userSource

import (
	"github.com/AitorGuerrero/UserGo/user"
	"errors"
//	"fmt"
)

type UserSource struct {
	collection []user.User
}

func (us *UserSource) Add (u user.User) error {
	us.collection = append(us.collection, u)

	return nil
}

func (us UserSource) ById (i user.Id) (user.User, error) {
	for _, u := range us.collection {
		if u.Id() == i {
			return u, nil
		}
	}

	return user.User{}, errors.New("Not found user "+string(i))
}
