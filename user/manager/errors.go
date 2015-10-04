package manager

import "github.com/AitorGuerrero/UserGo/user"

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

type ManagerDoesNotOwnTheNamespace struct {
	Manager Manager
	Namespace user.Namespace
}
func (ManagerDoesNotOwnTheNamespace) Error() string {
	return "Manager does not own the namespace"
}
