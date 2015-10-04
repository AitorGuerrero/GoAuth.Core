package manager

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type ExistentNamespaceError struct {}
func (ExistentNamespaceError) Error () string {
	return "A manager with that namespace exists"
}

type DuplicatedIdError struct {}
func (DuplicatedIdError) Error () string {
	return "A manager with that ID exists"
}

type Source interface {
	Add (*Manager) error
	Get (i user.Id) (*Manager, error)
	ExistsWithNamespace(user.Namespace) bool
}
