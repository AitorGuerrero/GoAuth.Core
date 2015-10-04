package user

import (
	t "testing"
)

func TestWhenCreatedNewShouldBeEmpty(t *t.T) {
	tk := Token{}

	if !tk.IsEmpty() {
		t.Error(tk)
	}
}
