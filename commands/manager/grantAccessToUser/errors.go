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
