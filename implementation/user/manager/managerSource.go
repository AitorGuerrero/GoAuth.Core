package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"

	"errors"
)

type Source struct {
	Collection map[string]*manager.Manager
}

func (s *Source) Add (m *manager.Manager) (err error) {
	if s.ExistsWithNamespace(m.Namespace) {
		err = manager.ExistentNamespaceError{}
		return
	}
	if s.Collection[string(m.Id())] != nil {
		err = manager.DuplicatedIdError{}
		return
	}
	s.Collection[string(m.Id())] = m

	return
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
