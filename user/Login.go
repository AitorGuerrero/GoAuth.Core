package user

type Login struct {
	Validator SignInValidator
	UserSource UserSource
}

func (l Login) Try(uid Id, p Passkey) (u User, err error) {
	u, err = l.UserSource.Get(uid)
	if(nil != err) {
		return
	}
	err = l.Validator.Validate(u, p)
	if nil != err {
		return
	}
	if (false == u.HasToken()) {
		u.GenerateToken()
	}

	return
}
