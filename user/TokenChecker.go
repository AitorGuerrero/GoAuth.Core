package user

import (
	"errors"
)

type TokenChecker struct {
	TokenSource TokenSource
}

func (tc TokenChecker) Check(t Token) error {
	rt, err := tc.TokenSource.ByUser(t.User)
	if (nil != err || string(rt.Code) != string(t.Code)) {
		return errors.New("Incorrect Token")
	}

	return nil
}
