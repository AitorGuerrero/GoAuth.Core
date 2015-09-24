package userSource

import (
	t "testing"
	"github.com/AitorGuerrero/UserGo/user"
)

var identifier = user.Identifier("Pis");
var s = New();
var u = user.New(identifier);

func TestAddingAUserShouldStoreInCollection(t *t.T) {
	s.Add(u);
	if 1 != len(s.collection) {
		t.Error("Should add the user to the collection")
	}
	if u != s.collection[0] {
		t.Error("Should save the same element")
	}
}

func TestShouldBeAbleToRetrieveAStoredUerByItsIdentifier(t *t.T) {
	s.Add(u);
	ru, error := s.ById(identifier);
	if error != nil {
		t.Error("The user should be in the source")
	}
	if ru != u {
		t.Error("The retrieved user should be the same as stored")
	}
}
