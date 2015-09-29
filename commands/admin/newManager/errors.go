package newManager

type EmptyPasskeyError struct {}

func (EmptyPasskeyError) Error() string {
	return "Passkey should not be empty"
}

type ExistingNamespaceError struct {}

func (ExistingNamespaceError) Error() string {
	return "A manager with that namespace exists"
}

type IncorrectIdError struct {}
func (IncorrectIdError) Error() string {
	return "Incorrect UUID id"
}
