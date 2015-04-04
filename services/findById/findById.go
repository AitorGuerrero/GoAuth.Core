package findById

import (
	userRepo "github.com/AitorGuerrero/User/user/persistence"
	"github.com/AitorGuerrero/User/user"

	"errors"
)

type Response struct {
	Id string
	Name string
	Email string
}

func Service(id string, ur userRepo.UserRepo) (Response, error) {
	u, error := ur.Find(user.Id(id))
	if error != nil {
		return Response{}, errors.New("A user with that id does not exists")
	}
	return Response{string(u.Id()), u.Name(), u.Email()}, nil
}
