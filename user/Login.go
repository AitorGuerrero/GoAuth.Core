package user

import "errors"

type Login struct {
	Validator SignInValidator
	UserSource Source
}

func (l Login) Try(uid Id, p string, n Namespace) (uo *User, err error) {
	// Demo of functional programming
	err = l.getUser(uid, func(u *User) {
		uo = u
		err = l.ifHasAccessToNamespace(*uo, n, func() {
			err = l.ifValidates(*uo, p, func() {
				l.ifHasNotToken(*uo, func() {
					uo.GenerateToken()
				})
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

func (l Login) ifValidates(u User, p string, fn func()) (err error) {
	err = l.Validator.Validate(u, p)
	if nil == err {
		fn()
	}

	return
}

func (l Login) getUser(uid Id, fn func(u *User)) (err error) {
	u, err := l.UserSource.Get(uid)
	if(nil == err) {
		fn(u)
	}

	return
}

func (l Login) ifHasAccessToNamespace(u User, n Namespace, fn func() ) (err error) {
	if false == u.hasAccessTo(n) {
		err = errors.New("User does not have access to the namespace")
		return
	}
	fn()

	return
}
