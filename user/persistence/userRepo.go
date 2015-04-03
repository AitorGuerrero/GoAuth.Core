package persistence

import(
	"github.com/AitorGuerrero/User/user"

	"code.google.com/p/go-uuid/uuid"
)

type UserRepo interface {
	Persist (u interface{})
	FindCountById(id uuid.UUID) int
	Find(id uuid.UUID) user.SerializedUser
}
