package user

type Login struct {
	Validator SignInValidator
	TokensSource TokenSource
}

func (l Login) Try(u User, p Passkey) (Token, error) {
	t := Token{}
	err := l.Validator.Validate(u, p)
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
