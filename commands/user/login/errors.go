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

type IncorrectNamespaceError struct {
	OriginalError error
}

func (e IncorrectNamespaceError) Error () string {
	return "User does not have access to the namespace"
}
