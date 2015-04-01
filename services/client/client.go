package client

import (
	commonClient "github.com/AitorGuerrero/BadassCity/common/services/client"
	"github.com/AitorGuerrero/BadassCity/user/config/devConfig"

	"github.com/koding/kite"
)

func Get() *kite.Client {
	config := devConfig.Get()
	return commonClient.Get(config.ServiceClientName, config.ServiceClientVersion, config.ServiceClientPort)
}
