package config

import (
	"github.com/AitorGuerrero/User/persistence/mongoDB"
)

func Get() Config {
	return Config {
		SqlDbConfig {
			Name: "databaseName",
			UserName: "username",
			Password: "password",
			Host: "host",
		},
		mongoDB.MongoDBConfig {
			Server: "mongodb://user:password@server:port/database",
		},
		KiteServiceConfig {
			Name: "BaddassCity.user",
			Version: "0.0.0",
			Port: 3635,
		},
	}
}
