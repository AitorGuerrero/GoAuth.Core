package user

import (
	"errors"
)

type TokenChecker struct {
}

func (tc TokenChecker) Check(u User, t Token, n Namespace) (err error) {
	if !u.HasToken() || !u.token.IsSame(t) || !u.hasAccessTo(n) {
		err = errors.New("Incorrect Token")
	}

	return err
}
