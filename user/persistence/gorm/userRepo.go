package gorm

import (
	"github.com/AitorGuerrero/User/user"

	"code.google.com/p/go-uuid/uuid"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	gm *gorm.DB
}

func Get(gm *gorm.DB) userRepo {
	aRepo := userRepo {
		gm: gm,
	}

	return aRepo
}

func (aRepo userRepo) FindCountById(id uuid.UUID) int {
	var count int
	users := []user.SerializedUser{}
	aRepo.gm.Where("id = ?", string(id)).Find(&users).Count(&count)
	return count
}

func (r userRepo) Persist (u interface{}) {
	r.gm.Create(u)
}
