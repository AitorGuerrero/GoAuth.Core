package login

import (
	"github.com/AitorGuerrero/UserGo/user"
	testUser "github.com/AitorGuerrero/UserGo/adapters/user"

	t "testing"
)

var sp = "passkey"
var uid = user.ParseId("f07e17d3-d688-4975-a080-dbdf16c504fa")
var n = user.Namespace("test/namespace")
var e user.Encryptor
var pe user.PasskeyEncryptor
var u user.User
var v user.SignInValidator
var l user.Login
var s testUser.Source
var com Command
var req Request

func beforeEach() {
	e = testUser.Encryptor{}
	pe = user.PasskeyEncryptor{e}
	p := pe.Encrypt(uid, sp)
	u = user.New(uid, p)

	u.GrantAccessTo(n)
	v = user.SignInValidator{pe}
	s = testUser.NewSource()
	s.Add(&u)
	l = user.Login{v, &s}
	com = Command{l, &s}
	req = Request{uid.Serialize(), sp, string(n)}
}

func TestIfUserDoesNotExistShouldReturnError(t *t.T) {
	beforeEach()
	req.Id = "ddf80715-37b0-4891-80b1-51e78e07b63f"
	_, err := com.Execute(req)
	if _, ok := err.(UserDoesNotExist); !ok {
		t.Error(err)
	}
}

func TestIfPasskeyIsNotCorrectShouldReturnError(t *t.T) {
	beforeEach()
	req.Passkey = "incorrectPasskey"
	_, err := com.Execute(req)
	if _, ok := err.(IncorrectPasskeyError); !ok {
		t.Error(err)
	}
}

func TestIfUserDoNotHaveAccessToNamespaceShouldReturnError (t *t.T) {
	beforeEach()
	req.Namespace = "incorrect/namespace"
	_, err := com.Execute(req)
	if _, ok := err.(IncorrectNamespaceError); !ok {
		t.Error(err)
	}
}
