package user

//import (
//	"code.google.com/p/go-uuid/uuid"
//)

//type authHash string
//

//type UserSource interface {
//	ById(Id) User
////	Add(User)
//}

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

func New (id Id, passkey cryptedPasskey) User {
	return User{id, passkey}
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

//func (aUser user) Id() Id {
//	return aUser.id
//}
//
//func (aUser user) Name() string {
//	return aUser.name
//}
//
//func (aUser user) Email() string {
//	return aUser.email
//}
//
//func generateAuthHash (name string, password string) authHash {
//	aAuthHash := authHash(name + "+" + password)
//	return aAuthHash
//}
