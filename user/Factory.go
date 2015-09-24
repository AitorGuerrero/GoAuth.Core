package user

type Factory struct {
	Encryptor PasskeyEncryptor
}

func (f Factory) Make (id Id, passkey Passkey) User {
	return User{id, f.Encryptor.Encrypt(id, passkey)};
}
