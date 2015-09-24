package user

type UserSource interface {
	Add (User) error
	ById(Id) (User, error)
}

type Id string
type Passkey string
type cryptedPasskey string

type User struct {
	id Id
	passkey cryptedPasskey
}

func (u User) Id() Id {
	return u.id
}

func (u User) EncryptedPasskey() cryptedPasskey {
	return u.passkey
}

type Encryptor interface {
	Encrypt(string) string
}

type PasskeyEncryptor struct {
	Encryptor Encryptor
}

func (pe PasskeyEncryptor) Encrypt(id Id, passkey Passkey) cryptedPasskey {
	return cryptedPasskey(pe.Encryptor.Encrypt(string(id) + string(passkey)))
}

type Factory struct {
	Encryptor PasskeyEncryptor
}

func (f Factory) Make (id Id, passkey Passkey) User {
	return User{id, f.Encryptor.Encrypt(id, passkey)};
}

