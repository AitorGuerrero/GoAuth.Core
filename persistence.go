package user

import (
	"code.google.com/p/go-uuid/uuid"
)

type SerializedUser struct {
	Id string
	Name string
	Email string
	AuthHash string
}

func (aUser *user) serialize () interface{} {
	return &SerializedUser {
		Id: string(aUser.id),
		Name: aUser.name,
		Email: aUser.email,
		AuthHash: string(aUser.authHash),
	}
}

func (aUser *SerializedUser) unserialize () *user {
	return &user {
		id: uuid.UUID(aUser.Id),
		name: aUser.Name,
		email: aUser.Email,
		authHash: authHash(aUser.AuthHash),
	}
}
