package newManager

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic/services"
)

var id = "userIdentifier";
var passkey = "userPasskey";
var namespace = "userPasskey";

var managerSource = services.ManagerSource();

var com = Command{managerSource, services.UserPassKeyEncryptor()};

func TestShouldPersistInASource(t *t.T) {
	err := com.Execute(Request{id, passkey, namespace});
	if nil != err {
		t.Error(err)
	}
	manager, err := managerSource.ById(user.Id(id));
	if nil != err {
		t.Error(err)
	}
	if string(manager.Id()) != id {
		t.Error("the user stored is not the same")
	}
}
