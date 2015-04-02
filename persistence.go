package User

type SerializedUser struct {
	Id string
	Name string
	Email string
	AuthHash string
}

func (SerializedUser) TableName() string {
	return "users"
}

func (aUser *user) Serialize () SerializedUser {
	return SerializedUser {
		Id: string(aUser.id),
		Name: aUser.name,
		Email: aUser.email,
		AuthHash: string(aUser.authHash),
	}
}

func (aUser *SerializedUser) Unserialize () user {
	return user {
		id: id(aUser.Id),
		name: aUser.Name,
		email: aUser.Email,
		authHash: authHash(aUser.AuthHash),
	}
}
