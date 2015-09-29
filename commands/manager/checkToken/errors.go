package checkToken

type InvalidIdError struct {}
func (InvalidIdError) Error() string {
	return "Invalid user ID"
}
