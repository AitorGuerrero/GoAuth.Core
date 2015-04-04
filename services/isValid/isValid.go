package isValid

import (
	"github.com/AitorGuerrero/User/user"
	userRepo "github.com/AitorGuerrero/User/user/persistence"
)

func Service(id string, ur userRepo.UserRepo) (bool, error) {
	e := ur.Exists(user.Id(id))
	return e, nil
}
