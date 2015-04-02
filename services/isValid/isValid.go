package isValid

import (
	"github.com/AitorGuerrero/User/queries/existsUserWithId"
	"github.com/AitorGuerrero/User/persistence/userRepo"

	"code.google.com/p/go-uuid/uuid"
)

func Service(id string, ur userRepo.UserRepo) (bool, error) {
	q := existsUserWithId.New(ur)
	result := q.Execute(uuid.UUID(id))
	return result, nil
}
