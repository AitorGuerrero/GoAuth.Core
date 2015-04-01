package newUser

import (
	"github.com/AitorGuerrero/BadassCity/user"
)

func Service(name, email, password string) (interface{}, error) {
	aUser := user.New(name, email, password)
	result := aUser.Id()
	return result, nil
}
