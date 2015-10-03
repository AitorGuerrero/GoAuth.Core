package grantAccessToUser

type DomainError error

type UserDoesNotExist struct {
	DomainError
	Id string
}
func(e UserDoesNotExist) Error() string {
	return "User with Id \""+e.Id+"\" des not exists"
}

type ManagerDoesNotExist struct {
	DomainError
	Id string
}
func (e ManagerDoesNotExist) Error() string {
	return "Manager with Id "+e.Id+" does not exists"
}

type ManagerDoesNotOwnTheUser struct {
	DomainError
	UserId string
	ManagerId string
}
func (e ManagerDoesNotOwnTheUser) Error() string {
	return "Manager \""+e.ManagerId+"\" does not own the user \""+e.UserId+"\""
}

type ManagerDoesNotOwnTheNamespace struct {
	DomainError
	ManagerId string
	Namespace string
}
func (e ManagerDoesNotOwnTheNamespace) Error() string {
	return "Manager \""+e.ManagerId+"\" does not own the namespace \""+e.Namespace+"\""
}
