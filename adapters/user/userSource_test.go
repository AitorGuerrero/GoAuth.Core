package user

import (
	"github.com/AitorGuerrero/UserGo/user"

	t "testing"
)

var ut = user.UserSourceTests{
	func() user.Source {
		s := NewSource()
		return &s
	},
	func() {},
}

func TestUserSource(t *t.T) {
	ut.Run(t)
}
