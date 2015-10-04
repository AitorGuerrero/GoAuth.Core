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
	token, err := c.getTokenFromUserIfCorrectLogin(
		user.Id(req.Id),
		req.Passkey,
		user.Namespace(req.Namespace),
	)
	if(nil != err) {
		return
	}
	res.SessionToken = token.Serialize()

	return
}

func (c Command) getTokenFromUserIfCorrectLogin(uid user.Id, up string, n user.Namespace) (tc user.Token, err error) {
	u, err := c.getUserIfCorrectLogin(uid, up, n)
	if nil != err {
		return
	}
 	tc = u.Token

	return
}

func (c Command) getUserIfCorrectLogin(uid user.Id, up string, n user.Namespace) (u *user.User, err error) {
	u, err = c.Login.Try(uid, up, n)
	if(nil != err) {
		return
	}

	return
}
