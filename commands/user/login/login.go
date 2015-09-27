package login

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Login user.Login
	UserSource user.UserSource
}

type Request struct {
	Id string
	Passkey string
}

type Response struct {
	SessionToken string
}

func (c Command) Execute(req Request) (res Response, err error) {
	tokenCode, err := c.getTokenCodeFromUserIfCorrectLogin(
		user.Id(req.Id),
		user.Passkey(req.Passkey),
	)
	res.SessionToken = string(tokenCode)

	return
}

func (c Command) getTokenCodeFromUserIfCorrectLogin(uid user.Id, up user.Passkey) (tc user.TokenCode, err error) {
	u, err := c.getUserIfCorrectLogin(uid, up)
	if nil != err {
		return
	}
 	tc = u.Token().Code

	return
}

func (c Command) getUserIfCorrectLogin(uid user.Id, up user.Passkey) (u user.User, err error) {
	u, err = c.tryLogin(uid, up)
	if(nil != err) {
		return
	}
	u, err = c.persistUser(u)

	return
}


func (c Command) persistUser(u user.User) (user.User, error) {
	c.UserSource.Persist(u)
	return u, nil
}

func (c Command) tryLogin(uid user.Id, p user.Passkey) (u user.User, err error) {
	u, err = c.getUser(uid)
	if(nil != err) {
		return
	}
	u, err = c.Login.Try(u, p)

	return
}

func (c Command) getUser(uid user.Id) (user.User, error) {
	return c.UserSource.Get(uid)
}

func (c Command) getTokenCodeFromUser(u user.User) user.TokenCode {
	return u.Token().Code
}
