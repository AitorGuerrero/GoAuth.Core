package newUser

import (
	"github.com/AitorGuerrero/User"
	"github.com/AitorGuerrero/User/persistence/userRepo"
)

func Service(name, email, password string, ur userRepo.UserRepo) (interface{}, error) {
	u := User.New(name, email, password)
	su := u.Serialize()
	ur.Persist(&su)
	result := u.Id()
	return result, nil
}
