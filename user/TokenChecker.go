package user

type TokenChecker struct {
}

type IncorrectTokenError struct {}
func (IncorrectTokenError) Error () string {
	return "There is no user with that ID"
}

func (tc TokenChecker) Check(u User, t Token, n Namespace) (err error) {
	if !u.HasToken() || !u.token.IsSame(t) || !u.hasAccessTo(n) {
		err = IncorrectTokenError{}
	}

	return err
}
