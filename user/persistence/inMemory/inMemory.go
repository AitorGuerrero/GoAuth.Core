package inMemory

import(
	"github.com/AitorGuerrero/User/user"
	"github.com/AitorGuerrero/User/user/persistence"
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

func (r UserRepo) Find(id user.Id) user.User {
	return r.c[string(id)].Unserialize()
}
