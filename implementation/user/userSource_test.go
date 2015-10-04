package user

import (
	"github.com/AitorGuerrero/UserGo/user"

	t "testing"
)

var e Encryptor
var pe user.PasskeyEncryptor
var i = user.ParseId("de1ec52a-21c6-491c-99f9-6f29bc1d565f")
var p user.Passkey
var u user.User
var s Source

func beforeEach() {
	e = Encryptor{}
	pe = user.PasskeyEncryptor{e}
	i = user.ParseId("de1ec52a-21c6-491c-99f9-6f29bc1d565f")
	p = pe.Encrypt(i, "passkey")
	u = user.New(i, p)
	s = NewSource()
}

func TestWhenAddingAUserIfIdExistsShouldReturnError(t *t.T) {
	beforeEach()
	u2 := user.New(i, p)
	s.Add(&u)
	err := s.Add(&u2)
	if _, ok := err.(user.DuplicatedIdError); !ok {
		t.Error(err)
	}
}

func TestGettingAUserIfNotExistsShouldReturnAError (t *t.T) {
	i := user.ParseId("8809a0cd-fd8c-4796-bb06-62d1134962f0")
	_, err := s.Get(i)
	if _, ok := err.(user.NotExistentUser); !ok {
		t.Error(err)
	}
}

func TestShouldReturnTheSameUserAsAdded (t *t.T) {
	s.Add(&u)
	ru, _ := s.Get(i)

	if !u.Id().Equal(ru.Id()) {
		t.Error("Id should be the same")
	}
	if !u.Token.IsSame(ru.Token) {
		t.Error("Token should be the same")
	}
}
