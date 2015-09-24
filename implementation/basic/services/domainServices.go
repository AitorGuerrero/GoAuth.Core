package services

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic"
)
var encrypter = basic.Encryptor{}
var userPassKeyEncryptor = user.PasskeyEncryptor{encrypter}
var facuserFactory = user.Factory{userPassKeyEncryptor}
var signInValidator = user.SignInValidator{userPassKeyEncryptor, &userSource}

func UserFactory() user.Factory {
	return facuserFactory
}

func UserPassKeyEncryptor() user.PasskeyEncryptor {
	return userPassKeyEncryptor
}

func SignInValidator() user.SignInValidator {
	return signInValidator
}
