package findById

import (
	userRepo "github.com/AitorGuerrero/User/user/persistence"
	"code.google.com/p/go-uuid/uuid"
)

type Response struct {
	Id string
	Name string
	Email string
}

func Service(id string, ur userRepo.UserRepo) (interface{}, error) {
	u := ur.Find(uuid.UUID(id))
	return Response{u.Id, u.Name, u.Email}, nil
}
