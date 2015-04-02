package userRepo

import(
	"code.google.com/p/go-uuid/uuid"
)

var impl UserRepo

type UserRepo interface {
	Persist (u interface{})
	FindCountById(id uuid.UUID) int
}

func Init (implementation UserRepo) {
	impl = implementation
}

func Get() UserRepo {
	return impl
}
