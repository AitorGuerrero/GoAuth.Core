package userServices

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic"
	"github.com/AitorGuerrero/UserGo/implementation/inMemory/userSource"
)
var enc = basic.Encryptor{}
var pe = user.PasskeyEncryptor{enc}

func Factory() user.Factory {
	return user.Factory{pe}
}

func PassKeyEncryptor() user.PasskeyEncryptor {
	return user.PasskeyEncryptor{enc}
}

func Source() user.UserSource {
	return &userSource.UserSource{};
}
