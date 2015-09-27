package user

import (
	"errors"
)

type TokenChecker struct {
}

func (tc TokenChecker) Check(u User, t Token) error {
	if (false == u.HasToken() || string(u.token.Code) != string(t.Code)) {
		return errors.New("Incorrect Token")
	}

	return nil
}
