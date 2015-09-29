package checkToken

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/implementation/services"
)

var com = Command{services.TokenChecker(), services.UserSource()}
var req = Request{}

func TestIfTheUserDoesNotExistsShouldReturnAnError (t *t.T) {
	req.Token = "PisToken"
	err := com.Execute(req)
	if _, ok := err.(InvalidIdError); !ok {
		t.Error()
	}
}
