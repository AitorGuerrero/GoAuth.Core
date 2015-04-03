package user

import (
	t "testing"
)

func TestWhenCreateAUserShouldCodeThePasswordCorrectly(t *t.T) {
	n := "aName"
	e := "aEmail"
	p := "aPassword"
	h := "aName+aPassword"
	u := New(n, e, p)
	if(string(u.authHash) != h) {
		t.Error("Hash is not well generated")
	}
}
