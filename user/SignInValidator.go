package user;

type SignInValidator struct {
	Encryptor PasskeyEncryptor
}

type IncorrectPasskeyError struct {}
func (IncorrectPasskeyError) Error () string {
	return "Incorrect passkey"
}

func (suv SignInValidator) Validate(u User, p string) error {
	if suv.Encryptor.Encrypt(u.Id(), p) != u.Passkey() {
		return IncorrectPasskeyError{}
	}

	return nil
}
