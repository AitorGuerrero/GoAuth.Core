package newUser

import (
	"github.com/AitorGuerrero/User"
)

func Service(name, email, password string) (interface{}, error) {
	aUser := User.New(name, email, password)
	result := aUser.Id()
	return result, nil
}
