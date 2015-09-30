package newUser

type ManagerDoesNotExistError struct {}
func (ManagerDoesNotExistError) Error() string {
	return "Managerwith that ID does not exists"
}
