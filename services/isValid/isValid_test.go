package isValid

import (
	t "testing"
	"github.com/AitorGuerrero/User/user/persistence/inMemory"
	"github.com/AitorGuerrero/User/user"
)

func TestWhenTheUserDoesNotExistsShouldShouldReturnFalse(t *t.T) {
	r := inMemory.New()
	id := "sadasdasdasd"
	exists, _ := Service(id, r)
	if exists {
		t.Error("User should not exists")
	}
}

func TestWhenUSerExistsShouldReturnTrue(t *t.T) {
	r := inMemory.New()
	u := user.New("a", "b", "c")
	r.Persist(u)
	exists, _ := Service(string(u.Id()), r)
	if !exists {
		t.Error("User should exists")
	}
}
