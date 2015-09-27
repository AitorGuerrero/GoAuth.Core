package user

type UserSource interface {
	Add (User) error
	Persist (User) error
	Get(Id) (User, error)
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

