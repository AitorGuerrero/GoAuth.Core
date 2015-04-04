package findById

import (
	t "testing"
	"github.com/AitorGuerrero/User/user/persistence/inMemory"
	"github.com/AitorGuerrero/User/user"
)

func TestWhenTheUserDoesNotExistsShouldThrowAnError(t *t.T) {
	r := inMemory.New()
	id := "sadasdasdasd"
	expectedError := "A user with that id does not exists"
	_, error := Service(id, r)
	if error.Error() != expectedError {
		t.Error("User should not exists:" + error.Error())
	}
}

func TestWhenUSerExistsShouldReturnTheUserData(t *t.T) {
	r := inMemory.New()
	u := user.New("a", "b", "c")
	r.Persist(u)
	ud, _ := Service(string(u.Id()), r)
	if ud.Id != string(u.Id()) {
		t.Error("Incorrect id")
	}
}
