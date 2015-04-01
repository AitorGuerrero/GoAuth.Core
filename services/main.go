package main

import (
	"github.com/AitorGuerrero/User"
	userRepo "github.com/AitorGuerrero/User/persistence/user/gorm"
	newUserService "github.com/AitorGuerrero/User/services/newUser/kite"
	isValidService "github.com/AitorGuerrero/User/services/isValid/kite"
	"github.com/AitorGuerrero/User/config"

	"github.com/koding/kite"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func main() {
	aConfig := config.Get()
	db, _ = initPersistence(aConfig.SqlDbConfig)
	initServices(aConfig.KiteServiceConfig)
	setRepos()
}

func initPersistence(dbConfig config.SqlDbConfig) (*gorm.DB, error) {
	db, _ := gorm.Open("mysql", dbConfig.UserName + ":" + dbConfig.Password + "@/" + dbConfig.Name + "?charset=utf8&parseTime=True&loc=Local")
	db.DB()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return &db, nil
}

func setRepos() {
	aRepo := userRepo.Get(db)
	User.SetRepo(&aRepo)
}

func initServices(kiteConfig config.KiteServiceConfig) {
	k := kite.New(kiteConfig.Name, kiteConfig.Version)
	newUserService.AddService(k)
	isValidService.AddService(k)

	k.Config.Port = kiteConfig.Port
	k.Run()
}
