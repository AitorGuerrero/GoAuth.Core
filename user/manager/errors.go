package manager

import "github.com/AitorGuerrero/UserGo/user"

type ExistentNamespaceError struct {}
func (ExistentNamespaceError) Error () string {
	return "A manager with that namespace exists"
}

type DuplicatedIdError struct {}
func (DuplicatedIdError) Error () string {
	return "A manager with that ID exists"
}

type NotExistsError struct {
	Id user.Id
}
func (NotExistsError) Error () string {
	return "A manager with that ID does not exists"
}

type DoesNotOwnTheUser struct {
	Manager *Manager
	User *user.User
}
func (DoesNotOwnTheUser) Error () string {
	return "The manager does not own the user"
}
