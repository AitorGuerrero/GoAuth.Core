package devConfig

import (
	"github.com/AitorGuerrero/BadassCity/user/config"
)

func Get() config.Config {
	return config.Config {
		DbName: "badass-city",
		DbUserName: "root",
		DbPassword: "devpassword",
		ServiceName: "BaddassCity.user",
		ServiceVersion: "0.0.1",
		ServicePort: 3635,
		ServiceClientName: "BaddassCity.user.client",
		ServiceClientVersion: "0.0.1",
		ServiceClientPort: 3636,
	}
}
