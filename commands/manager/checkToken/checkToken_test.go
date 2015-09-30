package checkToken

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/services"

	t "testing"
)

var userSource = services.UserSource()
var userFactory = services.UserFactory()
var com = Command{services.TokenChecker(), userSource}
var req = Request{}

func TestIfTheUserDoesNotExistsShouldReturnAnError (t *t.T) {
	req.Token = "PisToken"
	err := com.Execute(req)
	if _, ok := err.(InvalidIdError); !ok {
		t.Error()
	}
}

func TestIfTheTokenIsInvalidShouldReturnAnError (t *t.T) {
	pid, _ := user.ParseId("dabfe523-fae0-4a1c-8923-5e51ffeb3e91")
	u := userFactory.Make(pid, "passkey")
	u.GenerateToken()
	userSource.Add(&u)
	req.Token = "PisToken"
	req.UserId = "dabfe523-fae0-4a1c-8923-5e51ffeb3e91"
	err := com.Execute(req)
	if _, ok := err.(IncorrectTokenError); !ok {
		t.Error(err)
	}
}
