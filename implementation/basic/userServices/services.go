package userServices

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic"
	"github.com/AitorGuerrero/UserGo/implementation/inMemory/userSource"
)
var enc = basic.Encryptor{}
var pe = user.PasskeyEncryptor{enc}
var fac = user.Factory{pe}
var src = userSource.UserSource{}

func Factory() user.Factory {
	return fac
}

func PassKeyEncryptor() user.PasskeyEncryptor {
	return pe
}

func Source() user.UserSource {
	return &src;
}
