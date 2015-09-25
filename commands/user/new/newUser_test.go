package new

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/services"
)

var id = "userIdentifier";
var passkey = "userPasskey";

var userSource = services.UserSource();

var com = Command{userSource, services.UserFactory()};

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

