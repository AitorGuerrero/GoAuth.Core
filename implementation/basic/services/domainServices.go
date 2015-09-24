package services

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/session"
	"github.com/AitorGuerrero/UserGo/implementation/basic"
)
var encrypter = basic.Encryptor{}
var userPassKeyEncryptor = user.PasskeyEncryptor{encrypter}
var userFactory = user.Factory{userPassKeyEncryptor}
var signInValidator = user.SignInValidator{userPassKeyEncryptor, &userSource}
var login = session.Login{&userSource, signInValidator, &tokenSource}

func UserFactory() user.Factory {
	return userFactory
}

func UserPassKeyEncryptor() user.PasskeyEncryptor {
	return userPassKeyEncryptor
}

func SignInValidator() user.SignInValidator {
	return signInValidator
}

func UserLogin() session.Login {
	return login
}
