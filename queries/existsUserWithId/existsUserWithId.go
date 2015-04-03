package existsUserWithId

import (
	userRepo "github.com/AitorGuerrero/User/user/persistence"

	"code.google.com/p/go-uuid/uuid"
)

type query struct {
	ur userRepo.UserRepo
}

func New(ur userRepo.UserRepo) *query {
	return &(query{ur})
}

func (q query) Execute(anId uuid.UUID) bool {
	return q.ur.FindCountById(anId) > 0
}