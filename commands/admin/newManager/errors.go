package newManager

type EmptyPasskeyError struct {}

func (e EmptyPasskeyError) Error() string {
	return "Passkey should not be empty"
}

type ExistingNamespaceError struct {}

func (e ExistingNamespaceError) Error() string {
	return "A manager with that namespace exists"
}
