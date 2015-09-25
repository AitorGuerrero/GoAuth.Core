package user;

import(
	"errors"
)

type SignInValidator struct {
	Encryptor PasskeyEncryptor
}

func (suv SignInValidator) Validate(u User, p Passkey) error {
	if suv.Encryptor.Encrypt(u.Id(), p) != u.EncryptedPasskey() {
		return errors.New("The passkey doesn't match")
	}

	return nil
}
