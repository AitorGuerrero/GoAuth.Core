package gorm

import (
	"github.com/AitorGuerrero/User"

	"code.google.com/p/go-uuid/uuid"
	"github.com/jinzhu/gorm"
)

type repo struct {
	gm *gorm.DB
}

func Get(gm *gorm.DB) repo {
	aRepo := repo {
		gm: gm,
	}

	return aRepo
}

func (aRepo *repo) FindCountById(id uuid.UUID) int {
	var count int
	users := []User.SerializedUser{}
	aRepo.gm.Where("id = ?", string(id)).Find(&users).Count(&count)
	return count
}
