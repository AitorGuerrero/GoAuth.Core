package login

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/session"
	"github.com/AitorGuerrero/UserGo/implementation/basic/services"
)

var source = services.UserSource()
var id = "userIdentifier"
var passkey = "passkey"
var fac = services.UserFactory()
var tokenSource = session.TokenSource{};
var c = Command{source, services.SignInValidator(), &tokenSource}
var u = fac.Make(user.Id(id), user.Passkey(passkey))

func TestIfTheUserDoNotExistsShouldThrowAnError(t *t.T) {
	invalidUserIdentifier := "invalidUserIdentifier"
	r := Request{Id: invalidUserIdentifier}

	_, err := c.Execute(r)

	if nil == err {
		t.Error("When the user dont exists Should throw an error")
	}
}

func TestIfThePassKeyDoesNotMatchShouldThrowAnError(t *t.T) {
	source.Add(u)
	r := Request{id, "InvalidPasskey"}
	_, err := c.Execute(r)
	if nil == err {
		t.Error("When the passkey is invalid Should throw an error")
	}
}

func TestShouldReturnASessionToken(t * t.T) {
	source.Add(u)
	req := Request{id, passkey}
	res, _ := c.Execute(req)

	if (res.SessionToken == "") {
		t.Error("Should return a token")
	}

	_, err := tokenSource.ByUser(u)

	if (nil != err) {
		t.Error("Should store the token")
	}
}
