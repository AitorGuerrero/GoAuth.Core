package newUser

import (
	t "testing"
	"github.com/AitorGuerrero/User/user/persistence/inMemory"

	"code.google.com/p/go-uuid/uuid"
)

type user interface {
	Id() uuid.UUID
}

func TestShouldCreateNewUser(t *t.T) {
	n := "aName"
	e := "aEmail"
	p := "aPassword"
	r := inMemory.New()
	userId, _ := Service(n, e, p, r)
	resultingUser := r.Find(uuid.UUID(userId))
	if userId != string(resultingUser.Id()) {
		t.Error("User Id dont match" + string(userId))
	}
}
