package user

import (
	"code.google.com/p/go-uuid/uuid"
	"errors"
)

type Source interface {
	Add (User) error
	Persist (User) error
	Get(Id) (User, error)
}

type Id uuid.UUID
type Passkey string
type cryptedPasskey string

type User struct {
	id Id
	passkey cryptedPasskey
	token Token
	namespaces []Namespace
}

func (u User) Id() Id {
	return u.id
}

func (i Id) Equal (i2 Id) bool {
	return string(i) == string(i2)
}

func (i Id) IsEmpty () bool {
	return "" == string(i)
}

func ParseId (s string) (i Id, err error) {
	i = Id(uuid.Parse(s));
	if nil == i {
		err = errors.New("Invalid Id")
	}

	return
}

func (u User) EncryptedPasskey() cryptedPasskey {
	return u.passkey
}

func (u *User) GrantAccessTo(n Namespace) {
	u.namespaces = append(u.namespaces, n)
}

func (u User) hasAccessTo(n Namespace) bool {
	for _, ns := range u.namespaces {
		if ns.Owns(n) {
			return true
		}
	}

	return false
}
