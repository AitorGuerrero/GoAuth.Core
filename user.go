package user

import (
	"code.google.com/p/go-uuid/uuid"
)

var aRepo repository

type authHash string

type user struct {
	id uuid.UUID
	name string
	email string
	authHash authHash
}

type repository interface {
	FindCountById(uuid.UUID) int
}

func New (name string, email string, password string) user {
	return user{
		name: name,
		email: email,
		authHash: generateAuthHash(name, password),
	}
}

func SetRepo(r repository) {
	aRepo = r
}

func (aUser user) Id() uuid.UUID {
	return aUser.id
}

func generateAuthHash (name string, password string) authHash {
	aAuthHash := authHash(name + password)
	return aAuthHash
}

func ExistsUserWithId (aId uuid.UUID) bool {
	return aRepo.FindCountById(aId) > 0
}
