package grantAccessToUser

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
	managerImplementation "github.com/AitorGuerrero/UserGo/implementation/user/manager"
	userImplementation "github.com/AitorGuerrero/UserGo/implementation/user"

	t "testing"
)

var managerId = "dabfe523-fae0-4a1c-8923-5e51ffeb3e91"
var parsedManagerId, _ = user.ParseId(managerId)
var userId = "bbc99515-779c-471a-a43b-350184dc2569"
//var parsedUserrId, _ = user.ParseId(userId)
var namespace = "/test/passkey"
var managerNamespace = "/test"
var managerSource *managerImplementation.Source
var userSource *userImplementation.Source
var factory = user.Factory{user.PasskeyEncryptor{userImplementation.Encryptor{}}}
var com Command
var req Request
var m *manager.Manager

func beforeEach() {{}
	managerSource = &managerImplementation.Source{map[string]*manager.Manager{}}
	userSource = &userImplementation.Source{map[string]*user.User{}}
	req = Request{managerId, userId, namespace}
	com = Command{userSource, managerSource}
	m = createManager()
}

func createManager() (m *manager.Manager) {
	m = &manager.Manager{
		factory.Make(parsedManagerId, "passkey"),
		user.Namespace(managerNamespace),
		map[string]*user.User{},
	}
	managerSource.Add(m)

	return
}

func TestWhenUserDoNotExistsShouldReturnAnError (t *t.T) {
	beforeEach()
	err := com.Execute(req)
	if _, ok := err.(UserDoesNotExist); !ok {
		t.Error(err)
	}
}
