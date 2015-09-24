package newUser

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic/services"
)

var id = "userIdentifier";
var passkey = "userPasskey";

var userSource = services.UserSource();

var com = Command{userSource, services.UserPassKeyEncryptor()};

func TestTheUserShouldHaveAnUniqueIdentifier(t *t.T) {
	err := com.Execute(Request{Passkey: passkey});
	if (err == nil) {
		t.Error("Should throw an error")
	}
}

func TestShouldPersistInASource(t *t.T) {
	err := com.Execute(Request{id, passkey});
	if nil != err {
		t.Error(err)
	}
	u, err := userSource.ById(user.Id(id));
	if nil != err {
		t.Error(err)
	}
	if string(u.Id()) != id {
		t.Error("the user stored is not the same")
	}
}

