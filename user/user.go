package user

type Passkey string

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

func (u User) HasAccessTo(n Namespace) bool {
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

func (u User) CheckToken(t Token) bool {
	if !u.HasToken() || !u.Token.IsSame(t) {
		return false
	}

	return true
}
