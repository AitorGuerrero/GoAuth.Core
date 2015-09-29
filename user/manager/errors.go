package manager

type ExistentNamespaceError struct {}
func (ExistentNamespaceError) Error () string {
	return "A manager with that namespace exists"
}

type DuplicatedIdError struct {}
func (DuplicatedIdError) Error () string {
	return "A manager with that ID exists"
}
