package user

type ManagerSource interface {
	Add(Manager) error
	ById (i Id) (Manager, error)
}

type Namespace string

type Manager struct {
	User
	Namespace Namespace
}
