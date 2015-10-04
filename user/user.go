package user

type Passkey string

type DuplicatedIdError struct {
	Id Id
}
func (DuplicatedIdError) Error() string {
	return "The id is in use"
}
type NotExistentUser struct {
	Id Id
}
func (NotExistentUser) Error() string {
	return "A user with that ID does not exist"
}

type Source interface {
	Add (*User) error
	Get(Id) (*User, error)
}

type User struct {
	id Id
	passkey Passkey
	Token Token
	namespaces []Namespace
}

func New (i Id, p Passkey) User {
	return User{
		id: i,
		passkey: p,
	}
}

func (u User) Id() Id {
	return u.id
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

func (u User) HasToken() bool {
	return !u.Token.IsEmpty()
}

func (u *User) GenerateToken() {
	u.Token = GenerateToken()
}
