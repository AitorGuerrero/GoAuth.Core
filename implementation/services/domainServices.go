package services

import (
	"github.com/AitorGuerrero/UserGo/user"
	implementationUser "github.com/AitorGuerrero/UserGo/implementation/user"
)
var encrypter = implementationUser.Encryptor{}
var userPassKeyEncryptor = user.PasskeyEncryptor{encrypter}
var userFactory = user.Factory{userPassKeyEncryptor}
var signInValidator = user.SignInValidator{userPassKeyEncryptor}
var login = user.Login{signInValidator, &tokenSource}
var tokenChecker = user.TokenChecker{&tokenSource}

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

func TokenChecker() user.TokenChecker {
	return tokenChecker
}
