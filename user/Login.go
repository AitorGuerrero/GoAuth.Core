package user

type Login struct {
	UserSource UserSource
	Validator SignInValidator
	TokensSource TokenSource
}

func (l Login) Try(id Id, p Passkey) (Token, error) {
	t := Token{}
	u, err := l.UserSource.Get(id)
	if nil != err {
		return t, err
	}
	err = l.Validator.Validate(u, p)
	if nil != err {
		return t, err
	}
	t, err = l.TokensSource.ByUser(u)
	if (nil != err) {
		t = GenerateNewToken(u)
		l.TokensSource.Add(t)
	}

	return t, nil
}
