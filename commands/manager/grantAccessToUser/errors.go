package grantAccessToUser

type UserDoesNotExist struct {
	Id string
	OrigError error
}
func(UserDoesNotExist) Error() string {
	return "User With that Id des not exists"
}

type ManagerDoesNotExist struct {
	Id string
	OrigError error
}
func (ManagerDoesNotExist) Error() string {
	return "Manager With that Id des not exists"
}

type ManagerDoesNotOwnTheUser struct {
	UserId string
	ManagerId string
	OrigError error
}
func (e ManagerDoesNotOwnTheUser) Error() string {
	return "Manager "+e.ManagerId+" does not own the user "+e.UserId
}

type ManagerDoesNotOwnTheNamespace struct {
	OrigErr error
}
func (ManagerDoesNotOwnTheNamespace) Error() string {
	return "Manager does not own the namespace"
}
