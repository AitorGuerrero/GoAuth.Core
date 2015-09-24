package inMemory

import (
	"github.com/AitorGuerrero/UserGo/user"

	"errors"
)

type ManagerSource struct {
	collection []user.Manager
}

func (c *ManagerSource) Add (u user.Manager) error {
	c.collection = append(c.collection, u)

	return nil
}

func (ms ManagerSource) ById (i user.Id) (user.Manager, error) {
	for _, m := range ms.collection {
		if m.Id() == i {
			return m, nil
		}
	}

	return user.Manager{}, errors.New("Not found user")
}
