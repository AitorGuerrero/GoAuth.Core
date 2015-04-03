package newUser

import (
	"github.com/AitorGuerrero/User/user"
	userRepo "github.com/AitorGuerrero/User/user/persistence"
)

func Service(name, email, password string, ur userRepo.UserRepo) (string, error) {
	u := user.New(name, email, password)
	ur.Persist(u)
	return string(u.Id()), nil
}
