package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Source interface {
	Add (*Manager) error
	Get (i user.Id) (*Manager, error)
	ExistsWithNamespace(user.Namespace) bool
}

type Manager struct {
	user.User
	Namespace user.Namespace
	Users map[string]*user.User
}

func (m *Manager) AddUser(u *user.User) {
	m.Users[string(u.Id())] = u
}

func (m Manager) GrantAccessToUser(u *user.User, n user.Namespace) (err error) {
	if false == m.ownsNamespace(n) {
		err = ManagerDoesNotOwnTheNamespace{m, n}
		return
	}
	if false == m.ownsUser(*u) {
		err = DoesNotOwnTheUser{&m, u}
		return
	}
	u.GrantAccessTo(n)

	return
}

func (m Manager) ownsNamespace(n user.Namespace) bool {
	return m.Namespace.Owns(n)
}

func (m Manager) ownsUser(u user.User) bool {
	return false == u.Id().IsEmpty() && m.Users[string(u.Id())] != nil
}
