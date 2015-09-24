package user;

import(
	"errors"
)

type SignInValidator struct {
	Encryptor PasskeyEncryptor
	Source UserSource
}

func (suv SignInValidator) Validate(i Id, p Passkey) error {
	u, err := suv.Source.ById(i)
	if (nil != err) {
		return errors.New("The user with the that Id doesn't exists")
	}
	if suv.Encryptor.Encrypt(i, p) != u.EncryptedPasskey() {
		return errors.New("The passkey doesn't match")
	}

	return nil
}
