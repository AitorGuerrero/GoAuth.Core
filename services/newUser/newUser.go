package newUser

import (
	User "github.com/AitorGuerrero/User/user"
	userRepo "github.com/AitorGuerrero/User/user/persistence"
)

func Service(name, email, password string, ur userRepo.UserRepo) (string, error) {
	u := User.New(name, email, password)
	ur.Persist(u)
	result := u.Id()
	return string(result), nil
}
