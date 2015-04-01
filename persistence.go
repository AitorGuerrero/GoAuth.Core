package User

type SerializedUser struct {
	Id string
	Name string
	Email string
	AuthHash string
}

func (aUser *user) Serialize () interface{} {
	return &SerializedUser {
		Id: string(aUser.id),
		Name: aUser.name,
		Email: aUser.email,
		AuthHash: string(aUser.authHash),
	}
}

func (aUser *SerializedUser) Unserialize () *user {
	return &user {
		id: id(aUser.Id),
		name: aUser.Name,
		email: aUser.Email,
		authHash: authHash(aUser.AuthHash),
	}
}
