package persistence

import(
	"github.com/AitorGuerrero/User/user"
)

type Serializable interface {
	Serialize () *user.SerializedUser
}
type UserRepo interface {
	Persist (u user.User)
	Find(id user.Id) user.User
}
