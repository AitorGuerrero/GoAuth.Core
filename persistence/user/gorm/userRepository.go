package gorm

import (
	gm "github.com/AitorGuerrero/BadassCity/persistence/gorm"
	"github.com/AitorGuerrero/BadassCity/user"

	"code.google.com/p/go-uuid/uuid"
)

type repo struct {

}

func Get() repo {
	aRepo := repo {

	}

	return aRepo
}

func (*repo) FindCountById(id uuid.UUID) int {
	var count int
	users := []user.SerializedUser{}
	gm.Manager().db.Where("id = ?", string(id)).Find(&users).Count(&count)
	return count
}
