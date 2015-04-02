package main

import (
	"github.com/AitorGuerrero/User/persistence/userRepo"
	userRepoGorm "github.com/AitorGuerrero/User/persistence/userRepo/gorm"
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
var ur userRepo.UserRepo

func main() {
	c := config.Get()
	db, _ = initPersistence(c.SqlDbConfig)
	setRepos()
	initServices(c.KiteServiceConfig)
}

func initPersistence(c config.SqlDbConfig) (*gorm.DB, error) {
	db, _ := gorm.Open("mysql", c.UserName + ":" + c.Password + "@/" + c.Name + "?charset=utf8&parseTime=True&loc=Local")
	db.DB()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return &db, nil
}

func setRepos() {
	userRepo.Init(userRepoGorm.Get(db))
	ur = userRepo.Get()
}

func initServices(c config.KiteServiceConfig) {
	k := kite.New(c.Name, c.Version)

	newUserService.AddService(k, ur)
	isValidService.AddService(k, ur)

	k.Config.Port = c.Port
	k.Run()
}
