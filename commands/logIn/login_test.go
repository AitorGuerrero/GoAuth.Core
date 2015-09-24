package login

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic/userServices"
)

var source = userServices.Source()
var id = user.Id("userIdentifier");
var passkey = user.Passkey("passkey")

var pe = userServices.PassKeyEncryptor()
var fac = userServices.Factory()

var c = Command{source, pe}
var u = fac.Make(id, passkey)

func TestIfTheUserDoNotExistsShouldThrowAnError(t *t.T) {
	invalidUserIdentifier := user.Id("invalidUserIdentifier")
	r := Request{Id: invalidUserIdentifier}

	_, err := c.Execute(r)

	if nil == err {
		t.Error("When the user dont exists Should throw an error")
	}
}

func TestIfThePassKeyDoesNotMatchShouldThrowAnError(t *t.T) {
	source.Add(u)
	r := Request{id, user.Passkey("InvalidPasskey")}
	_, err := c.Execute(r)
	if nil == err {
		t.Error("When the passkey is invalid Should throw an error")
	}
}

func TestShouldReturnASessionToken(t * t.T) {
	source.Add(u)
	req := Request{id, user.Passkey(passkey)}
	res, _ := c.Execute(req)

	if (res.SessionToken == "") {
		t.Error("Should return a token")
	}
}