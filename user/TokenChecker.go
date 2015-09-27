package user

import (
	"errors"
)

type TokenChecker struct {
	TokenSource TokenSource
}

func (tc TokenChecker) Check(u User, t Token) error {
	rt, err := tc.TokenSource.ByUser(u)
	if (nil != err || string(rt.Code) != string(t.Code)) {
		return errors.New("Incorrect Token")
	}

	return nil
}
