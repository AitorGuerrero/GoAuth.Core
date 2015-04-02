package User

import (
	"code.google.com/p/go-uuid/uuid"
)

type authHash string

type id uuid.UUID

type user struct {
	id id
	name string
	email string
	authHash authHash
}

type repository interface {
	FindCountById(id uuid.UUID) int
}

func New (name string, email string, password string) user {
	return user{
		id: id(uuid.NewUUID()),
		name: name,
		email: email,
		authHash: generateAuthHash(name, password),
	}
}

func (aUser user) Id() uuid.UUID {
	return uuid.UUID(aUser.id)
}

func generateAuthHash (name string, password string) authHash {
	aAuthHash := authHash(name + password)
	return aAuthHash
}
