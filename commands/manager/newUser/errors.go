package newUser

type ManagerDoesNotExistError struct {
	OriginalError error
}
func (ManagerDoesNotExistError) Error() string {
	return "Managerwith that ID does not exists"
}

type MalformedUserIdError struct {
	OriginalError error
}
func (MalformedUserIdError) Error() string {
	return "The user ID is not a valid UUID"
}

type DuplicatedIdError struct {
	OriginalError error
}
func (DuplicatedIdError) Error() string {
	return "The user ID is not a valid UUID"
}
