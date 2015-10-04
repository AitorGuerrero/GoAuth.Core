package user

type TokenChecker struct {
}

type IncorrectTokenError struct {}
func (IncorrectTokenError) Error () string {
	return "There is no user with that ID"
}

type AccessErrorToNamespace struct {}
func (AccessErrorToNamespace) Error() string {
	return "User do not have access to that namespace"
}

func (tc TokenChecker) Check(u User, t Token, n Namespace) (error) {
	if !u.HasToken() || !u.Token.IsSame(t) {
		return IncorrectTokenError{}
	}
	if !u.hasAccessTo(n) {
		return AccessErrorToNamespace{}
	}

	return nil
}
