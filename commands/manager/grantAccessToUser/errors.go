package grantAccessToUser

type UserDoesNotExist struct {
	Id string
	OrigError error
}
func(UserDoesNotExist) Error() string {
	return "User With that Id des not exists"
}
