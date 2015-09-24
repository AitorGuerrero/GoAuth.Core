package checkToken

import (
	t "testing"
//	"github.com/AitorGuerrero/commands/newUser"
//	"github.com/AitorGuerrero/UserGo/implementation/basic/userServices"
)

var com = Command{}
var req = Request{}
//var userFactory = userServices.Factory()
//var userSource = userServices.Source()

func TestIfTheTokenDoesNotExistsShouldReturnAnError (t *t.T) {
	req.Token = "PisToken"
	err := com.Execute(req)
	if nil == err {
		t.Error("Should return an error")
	}
}

//func TestIfUserDoNotOwnTheTokenShouldReturnAnError (t *t.T) {
//	newUserCommand = newUser.Command{}
//	newUserCommand.Execute(newUser.Request{})
//	newUserCommand.Execute(newUser.Request{})
//
//	err := com.Execute(req)
//
//	if (nil == err) {
//		t.Error("Should throw an error")
//	}
//}
