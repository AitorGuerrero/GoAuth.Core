package manager

import (
	"github.com/AitorGuerrero/UserGo/user"

	"errors"
)

type ManagerSource interface {
	Add(Manager) error
	ById (i user.Id) (Manager, error)
	Persist(Manager)
}

type Manager struct {
	user.User
	Namespace user.Namespace
	users map[user.Id]user.User
}

func (m *Manager) AddUser(u user.User) {
	m.users[u.Id()] = u
}

func (m Manager) GrantAccessToUser(u user.User, n user.Namespace) (err error) {
	if false == m.ownsNamespace(n) {
		err = errors.New("Manager do not own this namespace")
		return
	}
	if false == m.ownsUser(u) {
		err = errors.New("Manager do not own this user")
		return
	}
	u.GrantAccessTo(n)

	return
}

func (m Manager) ownsNamespace(n user.Namespace) bool {
	return m.Namespace.Owns(n)
}

func (m Manager) ownsUser(u user.User) bool {
	return false == u.Id().IsEmpty() && m.users[u.Id()].Id().Equal(u.Id())
}
