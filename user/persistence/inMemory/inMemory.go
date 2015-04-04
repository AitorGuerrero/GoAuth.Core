package inMemory

import(
	"github.com/AitorGuerrero/User/user"
	"github.com/AitorGuerrero/User/user/persistence"

	"errors"
)

type UserRepo struct {
	c map[string]*user.SerializedUser
}

func New () *UserRepo {
	return &UserRepo{
		map[string]*user.SerializedUser{},
	}
}

func (r UserRepo) Persist (u user.User) {
	r.c[string(u.Id())] = u.(persistence.Serializable).Serialize()
}

func (r UserRepo) Find(id user.Id) (user.User, error) {
	if !r.Exists(id) {
		return nil, errors.New("Error")
	}
	return r.c[string(id)].Unserialize(), nil
}

func (r UserRepo) Exists(id user.Id) bool {
	e := false;
	if r.c[string(id)] != nil {
		e = true
	}

	return e
}
