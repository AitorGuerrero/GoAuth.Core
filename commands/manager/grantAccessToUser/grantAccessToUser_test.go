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
var parsedUserId, _ = user.ParseId(userId)
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
	m.AddUser(createUser())

	return
}

func createUser() (*user.User) {
	u := factory.Make(parsedUserId, "passkey")
	userSource.Add(&u)

	return &u
}

func TestWhenUserDoNotExistsShouldReturnAnError (t *t.T) {
	beforeEach()
	req.UserId = "cbee4f97-68e5-4979-9ba8-7a692464b548"
	err := com.Execute(req)
	if _, ok := err.(UserDoesNotExist); !ok {
		t.Error(err)
	}
}

func TestWhenManagerDoesNotExistsShouldReturnAnError (t *t.T) {
	beforeEach()
	req.ManagerId = "cbee4f97-68e5-4979-9ba8-7a692464b548"
	err := com.Execute(req)
	if _, ok := err.(ManagerDoesNotExist); !ok {
		t.Error(err)
	}
}
