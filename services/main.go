package main

import (
	newUserService "github.com/AitorGuerrero/BadassCity/user/services/newUser/kite"
	isValidService "github.com/AitorGuerrero/BadassCity/user/services/isValid/kite"
	"github.com/AitorGuerrero/BadassCity/user/config/devConfig"
	"github.com/AitorGuerrero/BadassCity/user/config"
	"github.com/AitorGuerrero/BadassCity/persistence/gorm"

	"github.com/koding/kite"
)

func main() {
	config := devConfig.Get()
	initPersistence(config)
	initServices(config)
}

func initPersistence(config config.Config) {
	gorm.Init(
		config.DbUserName,
		config.DbPassword,
		config.DbName,
	)
}

func initServices(config config.Config) {
	k := kite.New(config.ServiceName, config.ServiceVersion)
	newUserService.AddService(k)
	isValidService.AddService(k)

	k.Config.Port = config.ServicePort
	k.Run()
}
