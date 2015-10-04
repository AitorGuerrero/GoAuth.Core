package user

type DuplicatedIdError struct {
	Id Id
}
func (DuplicatedIdError) Error() string {
	return "The id is in use"
}
type NotExistentUser struct {
	Id Id
}
func (NotExistentUser) Error() string {
	return "A user with that ID does not exist"
}

type Source interface {
	Add (*User) error
	Get(Id) (*User, error)
}
