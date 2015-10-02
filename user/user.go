package user

type Passkey string

type DuplicatedIdError struct {
	Id Id
}
func (DuplicatedIdError) Error() string {
	return "The id is in use"
}

type Source interface {
	Add (*User) error
	Get(Id) (*User, error)
}

type User struct {
	id Id
	passkey Passkey
	token Token
	namespaces []Namespace
}

func (u User) Passkey() Passkey {
	return u.passkey
}

func (u *User) GrantAccessTo(n Namespace) {
	u.namespaces = append(u.namespaces, n)
}

func (u User) hasAccessTo(n Namespace) bool {
	for _, ns := range u.namespaces {
		if ns.Owns(n) {
			return true
		}
	}

	return false
}
