package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"

	"errors"
)

type Source struct {
	collection []manager.Manager
}

func (c *Source) Add (u manager.Manager) error {
	c.collection = append(c.collection, u)

	return nil
}

func (s *Source) Add (m manager.Manager) (err error) {
	if s.collection[m.Id()].Id() == m.Id() {
		err = errors.New("Existing user")
		return
	}
	s.Persist(m)

	return
}

func (ms Source) ById (i user.Id) (manager.Manager, error) {
	for _, m := range ms.collection {
		if m.Id() == i {
			return m, nil
		}
	}

	return manager.Manager{}, errors.New("Not found user")
}

func (s *Source) Persist (m manager.Manager) error {
	s.collection[m.Id()] = m

	return nil
}
