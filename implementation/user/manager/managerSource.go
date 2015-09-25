package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"

	"errors"
)

type ManagerSource struct {
	collection []manager.Manager
}

func (c *ManagerSource) Add (u manager.Manager) error {
	c.collection = append(c.collection, u)

	return nil
}

func (ms ManagerSource) ById (i user.Id) (manager.Manager, error) {
	for _, m := range ms.collection {
		if m.Id() == i {
			return m, nil
		}
	}

	return manager.Manager{}, errors.New("Not found user")
}
