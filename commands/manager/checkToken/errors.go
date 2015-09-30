package checkToken

type InvalidIdError struct {}
func (InvalidIdError) Error() string {
	return "Invalid user ID"
}

type IncorrectTokenError struct {}
func (IncorrectTokenError) Error() string {
	return "Incorrect token for user"
}
