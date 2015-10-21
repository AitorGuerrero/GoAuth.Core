package user

import "code.google.com/p/go-uuid/uuid"

type PlainUser struct {
	Id string
	Passkey string
	Token string
	Namespaces []string
}

func (u User) Plain() PlainUser {
	return PlainUser{
		Id: u.id.serialize(),
		Passkey: string(u.passkey),
		Token: u.Token.Serialize(),
	}
}

func FromPlain(p PlainUser) User {
	return User{
		id: ParseId(p.Id),
		passkey: Passkey(p.Passkey),
		Token: Token(uuid.Parse(p.Token)),
	}
}
