package isValid

import (
	"github.com/AitorGuerrero/User"

	"code.google.com/p/go-uuid/uuid"
)

func Service(id string) (bool, error) {
	result := user.ExistsUserWithId(uuid.UUID(id))
	return result, nil
}
