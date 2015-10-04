package user

import (
	t "testing"
)

type fakeEncryptor struct {}

func (fakeEncryptor) Encrypt(i string) string {
	return i
}

var e fakeEncryptor
var pe PasskeyEncryptor
var i = ParseId("de1ec52a-21c6-491c-99f9-6f29bc1d565f")
var p Passkey
var u User
var s Source

type UserSourceTests struct {
	Source func() Source
	AfterEach func()
}

func (ut UserSourceTests) Run(t *t.T) {
	ut.TestWhenAddingAUserIfIdExistsShouldReturnError(t)
	ut.AfterEach()
	ut.TestGettingAUserIfNotExistsShouldReturnAError(t)
	ut.AfterEach()
	ut.TestShouldReturnTheSameUserAsAdded(t)
	ut.AfterEach()
}

func (ust UserSourceTests) beforeEach() {
	s = ust.Source()
	e = fakeEncryptor{}
	pe = PasskeyEncryptor{e}
	p = pe.Encrypt(i, "passkey")
	u = New(i, p)
}

func (ust UserSourceTests) TestWhenAddingAUserIfIdExistsShouldReturnError(t *t.T) { // TODO Should be private
	ust.beforeEach()
	u2 := New(i, p)
	s.Add(&u)
	err := s.Add(&u2)
	if _, ok := err.(DuplicatedIdError); !ok {
		t.Error(err)
	}
}

func (ust UserSourceTests) TestGettingAUserIfNotExistsShouldReturnAError (t *t.T) {
	ust.beforeEach()
	i := ParseId("8809a0cd-fd8c-4796-bb06-62d1134962f0")
	_, err := s.Get(i)
	if _, ok := err.(NotExistentUser); !ok {
		t.Error(err)
	}
}

func (ust UserSourceTests) TestShouldReturnTheSameUserAsAdded (t *t.T) {
	ust.beforeEach()
	s.Add(&u)
	ru, _ := s.Get(i)

	if(ru == nil) {
		t.Error("User has not been saved")
	} else {
		if !u.Id().Equal(ru.Id()) {
			t.Error("Id should be the same")
		}
		if !u.Token.IsSame(ru.Token) {
			t.Error("Token should be the same")
		}
	}
}
