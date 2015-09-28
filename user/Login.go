package user

type Login struct {
	Validator SignInValidator
	UserSource UserSource
}

func (l Login) Try(uid Id, p Passkey) (uo User, err error) {

	// Demo of functional programming
	err = l.getUser(uid, func(u User) {
			uo = u
			err = l.ifValidates(uo, p, func() {
					l.ifHasNotToken(uo, func() {
							uo.GenerateToken()
						})
				})
		})

	return
}

func (l Login) ifHasNotToken(u User, fn func()) {
	if (false == u.HasToken()) {
		fn()
	}
}

func (l Login) ifValidates(u User, p Passkey, fn func()) (err error) {
	err = l.Validator.Validate(u, p)
	if nil == err {
		fn()
	}

	return
}

func (l Login) getUser(uid Id, fn func(u User)) (err error) {
	u, err := l.UserSource.Get(uid)
	if(nil == err) {
		fn(u)
	}

	return
}
