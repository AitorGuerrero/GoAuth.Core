package newUser

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic/userServices"
//	"fmt"
)

var id = user.Id("userIdentifier");
var passkey = user.Passkey("userPasskey");

var usi = userServices.Source();
var pe = userServices.PassKeyEncryptor()

var r = Request{}
var c = Command{usi, pe};

func TestTheUserShouldHaveAnUniqueIdentifier(t *t.T) {
	r = Request{passkey: passkey}
	err := c.Execute(r);
	if (err == nil) {
		t.Error("Should throw an error")
	}
}

func TestShouldPersistInASource(t *t.T) {
	var r = Request{id, passkey}
	err := c.Execute(r);
	if nil != err {
		t.Error(err)
	}
	u, err := usi.ById(id);
	if nil != err {
		t.Error(err)
	}
	if u.Id() != id {
		t.Error("the user stored is not the same:")
	}
}
