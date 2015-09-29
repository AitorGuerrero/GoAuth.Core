package user

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
)

var identifier = user.Id("Pis");
var passkey = user.Passkey("passkey")
var s = New();

func TestAddingAUserShouldStoreInCollection(t *t.T) {
	u := user.New(identifier, passkey);
	s.Add(u);
	if 1 != len(s.collection) {
		t.Error("Should add the user to the collection")
	}
	if u != s.collection[0] {
		t.Error("Should save the same element")
	}
}

func TestShouldBeAbleToRetrieveAStoredUserByItsIdentifier(t *t.T) {
	u := user.New(identifier, passkey)
	s.Add(u);
	ru, error := s.Get(identifier);
	if error != nil {
		t.Error("The user should be in the source")
	}
	if ru != u {
		t.Error("The retrieved user should be the same as stored")
	}
}
