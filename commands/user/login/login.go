package login

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Login user.Login
	UserSource user.Source
}

type Request struct {
	Id string
	Passkey string
	Namespace string
}

type Response struct {
	SessionToken string
}

func (c Command) Execute(req Request) (res Response, err error) {
	tokenCode, err := c.getTokenCodeFromUserIfCorrectLogin(
		user.Id(req.Id),
		req.Passkey,
		user.Namespace(req.Namespace),
	)
	if(nil != err) {
		return
	}
	res.SessionToken = string(tokenCode)

	return
}

func (c Command) getTokenCodeFromUserIfCorrectLogin(uid user.Id, up string, n user.Namespace) (tc user.TokenCode, err error) {
	u, err := c.getUserIfCorrectLogin(uid, up, n)
	if nil != err {
		return
	}
 	tc = u.Token().Code

	return
}

func (c Command) getUserIfCorrectLogin(uid user.Id, up string, n user.Namespace) (u user.User, err error) {
	u, err = c.Login.Try(uid, up, n)
	if(nil != err) {
		return
	}
	err = c.UserSource.Persist(u)

	return
}
