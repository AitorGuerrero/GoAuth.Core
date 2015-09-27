package user

type Login struct {
	Validator SignInValidator
}

func (l Login) Try(u User, p Passkey) (User, error) {
	err := l.Validator.Validate(u, p)
	if nil != err {
		return User{}, err
	}
	if (false == u.HasToken()) {
		u.GenerateToken()
	}

	return u, nil
}
