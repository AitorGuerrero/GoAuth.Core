package login

type UserDoesNotExist struct {
	OriginalError error
	Id string
}
func (e UserDoesNotExist) Error () string {
	return "User with that id does not exist"
}

type IncorrectPasskeyError struct {
	OriginalError error
	Passkey string
}
func (e IncorrectPasskeyError) Error () string {
	return "Incorrect passkey"
}
