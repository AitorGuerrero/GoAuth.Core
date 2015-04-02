package newUser

import (
	User "github.com/AitorGuerrero/User/user"
	userRepo "github.com/AitorGuerrero/User/user/persistence"
)

func Service(name, email, password string, ur userRepo.UserRepo) (interface{}, error) {
	u := User.New(name, email, password)
	su := u.Serialize()
	ur.Persist(&su)
	result := u.Id()
	return result, nil
}
