package user

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type BSONUser struct {
	Id string
	Passkey string
	Token string
	Namespaces []string
}

func (u User) GetBSON() (bu interface{}, err error) {
	ns := []string{}
	for _, n := range u.namespaces {
		ns = append(ns, string(n))
	}
	bu = BSONUser{
		Id: u.id.Serialize(),
		Passkey: string(u.passkey),
		Token: u.Token.Serialize(),
		Namespaces: ns,
	}

	a := BSONUser{
		Id: u.id.Serialize(),
		Passkey: string(u.passkey),
		Token: u.Token.Serialize(),
		Namespaces: ns,
	}
	fmt.Printf("GetBSON ID: %s\n", a.Id)

	return
}

func (u *User) SetBSON(raw bson.Raw) (err error) {
	bu := BSONUser{}
	err = raw.Unmarshal(&bu)
	if (err != nil) {
		return
	}
	u.id = ParseId(bu.Id)
	u.passkey = Passkey(bu.Passkey)
	for _, n := range bu.Namespaces {
		u.namespaces = append(u.namespaces, Namespace(n))
	}

	return
}
