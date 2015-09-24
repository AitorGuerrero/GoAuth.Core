package checkToken

import (
	t "testing"
//	"github.com/AitorGuerrero/UserGo/user"
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
//	userId := "userId";
//	userSource.Add(userFactory.Make(user.Id("userId"), user.Passkey("userPasskey")))
//	userSource.Add(userFactory.Make(user.Id("secondUserId"), user.Passkey("secondUserPasskey")))
//	req.Token =
//	err := com.Execute(req)
//}
