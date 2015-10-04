package user

type IncorrectNamespaceError struct {
	User User
	Namespace Namespace
}
func (IncorrectNamespaceError) Error () string {
	return "User does not have access to the namespace"
}

type Login struct {
	Validator SignInValidator
	UserSource Source
}

func (l Login) Try(uid Id, p string, n Namespace) (u *User, err error) {

	if u, err = l.UserSource.Get(uid); err != nil {
		return
	}
	if !u.hasAccessTo(n) {
		err = IncorrectNamespaceError{*u, n}
		return
	}
	if err = l.Validator.Validate(*u, p); err != nil {
		return
	}
	if !u.HasToken() {
		u.GenerateToken()
	}

	return
}
