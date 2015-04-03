package newUser

import (
	t "testing"
	"github.com/AitorGuerrero/User/user/persistence/inMemory"
	"github.com/AitorGuerrero/User/user"
)

func TestShouldCreateNewUser(t *t.T) {
	n := "aName"
	e := "aEmail"
	p := "aPassword"
	r := inMemory.New()
	userId, _ := Service(n, e, p, r)
	resultingUser := r.Find(user.Id(userId))
	if userId != string(resultingUser.Id()) {
		t.Error("User Id dont match" + string(userId))
	}
}
