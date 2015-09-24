package login

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/inMemory/userSource"
//		"fmt"
)

var source = userSource.New();
var id = user.Id("userIdentifier");
var passkey = user.Passkey("passkey")
var r = Request{}
var c = Command{&source}
var u = user.New(id, passkey)

func TestIfTheUserDownNotExistsShouldThrowAnError(t *t.T) {
	invalidUserIdentifier := user.Id("invalidUserIdentifier")
	r := Request{id: invalidUserIdentifier}

	err := c.Execute(r)

	if nil == err {
		t.Error("When the user dont exists Should throw an error")
	}
}

func TestIfThePassKeyDoesNotMatchShouldThrowAnError(t *t.T) {
	source.Add(u)
	r := Request{id, user.Passkey("InvalidPasskey")}
	err := c.Execute(r)
	if nil == err {
		t.Error("When the passkey is invalid Should throw an error")
	}
}
