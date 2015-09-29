package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"

	"errors"
)

type Source struct {
	Collection map[string]*manager.Manager
}

func (s *Source) Add (m *manager.Manager) error {
	if s.ExistsWithNamespace(m.Namespace) {
		return manager.ExistentNamespaceError{}

	}
	if s.Collection[string(m.Id())] != nil {
		return manager.DuplicatedIdError{}
	}
	s.Collection[string(m.Id())] = m

	return nil
}

func (ms Source) Get (i user.Id) (*manager.Manager, error) {
	for _, m := range ms.Collection {
		if m.Id().Equal(i) {
			return m, nil
		}
	}

	return &manager.Manager{}, errors.New("Not found manager")
}

func (ms Source) ExistsWithNamespace(n user.Namespace) bool {
	for _, m := range ms.Collection {
		if m.Namespace.Equal(n) {
			return true
		}
	}

	return false
}
