package user

import (
	"gopkg.in/mgo.v2/bson"
	t "testing"
)

var id = "dabfe523-fae0-4a1c-8923-5e51ffeb3e91"
var uuidId = ParseId(id)
var passkey = Passkey("userPasskey")
var user = User{
	id: uuidId,
	passkey: passkey,
}

func TestWhendMarshalAndUnmarshalShouldHaveTheCoreectData(t *t.T) {
	m, err := bson.Marshal(user)
	u := User{}
	if err != nil {
		t.Error(err)
	}
	err = bson.Unmarshal(m, &u)
	if err != nil {
		t.Error(err)
	}
	if (!u.id.Equal(uuidId)) {
		t.Errorf(
			"Id is not the same.\n Expected: %s\n Actual: %s",
			uuidId.Serialize(),
			u.id.Serialize(),
		)
	}
	if (u.passkey != passkey) {
		t.Errorf(
			"Passkey is not the same.\n Expected: %s\n Actual: %s",
			string(passkey),
			string(u.passkey),
		)
	}
}
