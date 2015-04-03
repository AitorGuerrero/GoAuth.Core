package user

import (
	"code.google.com/p/go-uuid/uuid"
)

type authHash string

type Id uuid.UUID

type user struct {
	id Id
	name string
	email string
	authHash authHash
}

type User interface {
	Id() Id
}

type repository interface {
	FindCountById(id uuid.UUID) int
}

func New (name string, email string, password string) *user {
	return &user{
		id: Id(uuid.NewUUID()),
		name: name,
		email: email,
		authHash: generateAuthHash(name, password),
	}
}

func (aUser user) Id() Id {
	return aUser.id
}

func generateAuthHash (name string, password string) authHash {
	aAuthHash := authHash(name + password)
	return aAuthHash
}
