package newUser

type ManagerDoesNotExistError struct {}
func (ManagerDoesNotExistError) Error() string {
	return "Managerwith that ID does not exists"
}

type MalformedUserIdError struct {}
func (MalformedUserIdError) Error() string {
	return "The user ID is not a valid UUID"
}
