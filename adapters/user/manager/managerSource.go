package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
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

func (s Source) Get (i user.Id) (m *manager.Manager, err error) {
	m = s.Collection[string(i)]
	if m == nil {
		err = manager.NotExistsError{i}
	}

	return m, err
}

func (ms Source) ExistsWithNamespace(n user.Namespace) bool {
	for _, m := range ms.Collection {
		if m.Namespace.Equal(n) {
			return true
		}
	}

	return false
}
