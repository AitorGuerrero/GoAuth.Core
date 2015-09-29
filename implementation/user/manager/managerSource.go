package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"

	"errors"
)

type Source struct {
	Collection map[user.Id]manager.Manager
}

func (s *Source) Add (m manager.Manager) (err error) {
	if s.Collection[m.Id()].Id() == m.Id() {
		err = errors.New("Existing user")
		return
	}
	s.Persist(m)

	return
}

func (ms Source) ById (i user.Id) (manager.Manager, error) {
	for _, m := range ms.Collection {
		if m.Id() == i {
			return m, nil
		}
	}

	return manager.Manager{}, errors.New("Not found user")
}

func (s *Source) Persist (m manager.Manager) error {
	s.Collection[m.Id()] = m

	return nil
}
