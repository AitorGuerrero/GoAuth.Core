package login

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/inMemory/userSource"
	"github.com/AitorGuerrero/UserGo/implementation/basic"
//		"fmt"
)

var source = userSource.New();
var id = user.Id("userIdentifier");
var passkey = user.Passkey("passkey")
var enc = basic.Encryptor{}
var pe = user.PasskeyEncryptor{enc}
var c = Command{&source, pe}
var fac = user.Factory{pe}
var u = fac.Make(id, passkey)

func TestIfTheUserDoNotExistsShouldThrowAnError(t *t.T) {
	invalidUserIdentifier := user.Id("invalidUserIdentifier")
	r := Request{Id: invalidUserIdentifier}

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
