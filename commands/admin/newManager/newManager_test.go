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
	if _, ok := err.(IncorrectIdError); !ok{
		t.Error()
	}
}

func TestIfNamespaceIsRepeatedShouldReturnError(t *t.T) {
	com.Execute(Request{id, passkey, namespace})
	err := com.Execute(Request{
		"c6384d58-fd79-49cb-a111-b08564e4db26",
		passkey,
		namespace,
	})
	if _, ok := err.(ExistingNamespaceError); !ok {
		t.Error()
	}
}

func TestIfPasskeyIsBlankShouldReturnAError(t *t.T) {
	err := com.Execute(Request{id, "", namespace})
	if _, ok := err.(EmptyPasskeyError); !ok {
		t.Error()
	}
}

func TestShouldAddInASource(t *t.T) {
	com.Execute(Request{id, passkey, namespace})
	manager, err := managerSource.Get(uuidId)
	if nil != err {
		t.Error(err, uuidId, managerSource)
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
		t.Error()
	}
}
