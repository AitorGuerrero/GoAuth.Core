package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type ManagerSource interface {
	Add(Manager) error
	ById (i user.Id) (Manager, error)
}

type Namespace string

type Manager struct {
	user.User
	Namespace Namespace
}
