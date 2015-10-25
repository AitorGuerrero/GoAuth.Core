package newUser


import(
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
	managerImplementation "github.com/AitorGuerrero/UserGo/adapters/user/manager"
	userImplementation "github.com/AitorGuerrero/UserGo/adapters/user"

	t "testing"
)

var managerId = "dabfe523-fae0-4a1c-8923-5e51ffeb3e91"
var parsedManagerId = user.ParseId(managerId)
var userId = "bbc99515-779c-471a-a43b-350184dc2569"
var parsedUserId = user.ParseId(userId)
var passKey = "userPasskey"
var managerPasskey = "managerPasskey"
var namespace = "/test/passkey"
var com Command
var req Request
var managerSource *managerImplementation.Source
var userSource *userImplementation.Source
var factory = user.Factory{user.PasskeyEncryptor{userImplementation.Encryptor{}}}
var m *manager.Manager


func beforeEach() {{}
	managerSource = &managerImplementation.Source{map[string]*manager.Manager{}}
	userSource = &userImplementation.Source{map[string]*user.User{}}
	com = Command{managerSource, userSource, factory}
	req = Request{managerId, userId, passKey}
	m = createManager()
}

func createManager() (m *manager.Manager) {
	m = &manager.Manager{
		factory.Make(parsedManagerId, managerPasskey),
		user.Namespace(namespace),
		map[string]*user.User{},
	}
	managerSource.Add(m)

	return
}

func TestIfManagerDoesNotExistsShouldReturnError (t *t.T) {
	beforeEach()
	req.ManagerId = "2244d379-c4f3-4ba1-9ffe-d4e57b52377a"
	err := com.Execute(req)
	if _, ok := err.(ManagerDoesNotExistError); !ok {
		t.Error(err)
	}
}

func TestIUserIdIsMalformedShouldReturnError (t *t.T) {
	beforeEach()
	req.UserId = "malformedUserId"
	err := com.Execute(req)
	if _, ok := err.(MalformedUserIdError); !ok {
		t.Error(err)
	}
}

func TestIUserIdExistsShouldReturnError (t *t.T) {
	beforeEach()
	u := factory.Make(parsedUserId, "anotherPasskey")
	userSource.Add(&u)
	err := com.Execute(req)
	if _, ok := err.(DuplicatedIdError); !ok {
		t.Error(err)
	}
}
