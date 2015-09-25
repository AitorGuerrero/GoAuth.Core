package services

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/implementation/basic"
)
var encrypter = basic.Encryptor{}
var userPassKeyEncryptor = user.PasskeyEncryptor{encrypter}
var userFactory = user.Factory{userPassKeyEncryptor}
var signInValidator = user.SignInValidator{userPassKeyEncryptor}
var login = user.Login{&userSource, signInValidator, &tokenSource}

func UserFactory() user.Factory {
	return userFactory
}

func UserPassKeyEncryptor() user.PasskeyEncryptor {
	return userPassKeyEncryptor
}

func SignInValidator() user.SignInValidator {
	return signInValidator
}

func UserLogin() user.Login {
	return login
}
