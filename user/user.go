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

type User struct {
	id Id
}

func (u User) Id() Id {
	return u.id
}

func New (id Id) User {
	return User{id}
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
