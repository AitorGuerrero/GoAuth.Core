package checkToken

import (
	"github.com/AitorGuerrero/UserGo/user"
	userImplementation "github.com/AitorGuerrero/UserGo/implementation/user"

	t "testing"
)

var id = "dabfe523-fae0-4a1c-8923-5e51ffeb3e91"
var userSource userImplementation.Source
var userFactory user.Factory
var com Command
var req Request
var parsedId user.Id
var passkeyEncryptor = user.PasskeyEncryptor{userImplementation.Encryptor{}}
var tokenChecker = user.TokenChecker{}

func beforeEach() {
	parsedId, _ = user.ParseId(id)
	userSource = userImplementation.Source{map[string]*user.User{}}
	userFactory = user.Factory{passkeyEncryptor}
	com = Command{tokenChecker, &userSource}
	req = Request{}
}

func TestIfTheUserDoesNotExistsShouldReturnAnError (t *t.T) {
	beforeEach()
	req.Token = "PisToken"
	err := com.Execute(req)
	if _, ok := err.(InvalidIdError); !ok {
		t.Error(err)
	}
}

func TestIfTheTokenIsInvalidShouldReturnAnError (t *t.T) {
	beforeEach()
	createUser()
	req.Token = "PisToken"
	req.UserId = id
	err := com.Execute(req)
	if _, ok := err.(IncorrectTokenError); !ok {
		t.Error(err)
	}
}

func TestIfTheUserHasNotAccessToNamespaceHouldReturnError (t *t.T) {
	beforeEach()
	u := createUser()
	req.Token = string(u.Token().Code)
	req.UserId = id
	err := com.Execute(req)
	if _, ok := err.(AccessErrorToNamespace); !ok {
		t.Error(err)
	}
}

func createUser() (u user.User) {
	u = userFactory.Make(parsedId, "passkey")
	u.GenerateToken()
	userSource.Add(&u)

	return
}
