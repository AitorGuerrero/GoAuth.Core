package persistence

import(
	"code.google.com/p/go-uuid/uuid"
)

type UserRepo interface {
	Persist (u interface{})
	FindCountById(id uuid.UUID) int
}
