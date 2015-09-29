package newManager

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/services"
)

var id = "dabfe523-fae0-4a1c-8923-5e51ffeb3e91"
var uuidId, _ = user.ParseId(id)
var passkey = "userPasskey"
var namespace = "userPasskey"

var managerSource = services.ManagerSource()

var com = Command{managerSource, services.UserFactory()}

func TestIfTheIdIsAInvalidUuidShouldThrowAnError(t *t.T) {
	err := com.Execute(Request{"incorrectUUID", passkey, namespace})
	if nil == err {
		t.Error()
	}
}

func TestIfPasskeyIsBlankShouldReturnAError(t *t.T) {
	err := com.Execute(Request{id, "", namespace})
	if nil == err {
		t.Error()
	}
}

func TestShouldAddInASource(t *t.T) {
	com.Execute(Request{id, passkey, namespace})
	manager, err := managerSource.Get(uuidId)
	if nil != err {
		t.Error("Error loading the user:", err)
		return
	}
	if !manager.Id().Equal(uuidId) {
		t.Error("the user stored is not the same")
		return
	}
}

func TestShouldCreateAToken(t * t.T) {
	com.Execute(Request{id, passkey, namespace})
	manager, _ := managerSource.Get(uuidId)
	if !manager.HasToken() {
		t.Error("Should create a token")
	}
}
