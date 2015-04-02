package main

import (
	"github.com/AitorGuerrero/User/persistence/mongoDB"
	userRepo "github.com/AitorGuerrero/User/user/persistence"
	userRepoMgo "github.com/AitorGuerrero/User/user/persistence/mgo"
	newUserService "github.com/AitorGuerrero/User/services/newUser/kite"
	isValidService "github.com/AitorGuerrero/User/services/isValid/kite"
	"github.com/AitorGuerrero/User/config"

	"github.com/koding/kite"
)

var c *mongoDB.Connection
var ur userRepo.UserRepo

func main() {
	cfg := config.Get()
	c, _ = mongoDB.Connect(cfg.MongoDBConfig)
	ur = userRepoMgo.Get(c)
	initServices(cfg.KiteServiceConfig)
}

func initServices(c config.KiteServiceConfig) {
	k := kite.New(c.Name, c.Version)

	newUserService.AddService(k, ur)
	isValidService.AddService(k, ur)

	k.Config.Port = c.Port
	k.Run()
}
