package inMemory

import(
	"github.com/AitorGuerrero/User/user"

	"code.google.com/p/go-uuid/uuid"
)

type UserRepo struct {
	c map[string]*user.SerializedUser
}

func New () *UserRepo {
	return &UserRepo{
		map[string]*user.SerializedUser{},
	}
}

func (r UserRepo) Persist (u interface{}) {
	r.c[u.(*user.SerializedUser).Id] = u.(*user.SerializedUser)
}

func (r UserRepo) FindCountById(id uuid.UUID) int {
	return 1
}

func (r UserRepo) Find(id uuid.UUID) *user.SerializedUser {
	return r.c[string(id)]
}
