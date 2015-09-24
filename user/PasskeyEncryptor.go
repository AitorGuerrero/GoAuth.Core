package user

type Encryptor interface {
	Encrypt(string) string
}

type PasskeyEncryptor struct {
	Encryptor Encryptor
}

func (pe PasskeyEncryptor) Encrypt(id Id, passkey Passkey) cryptedPasskey {
	return cryptedPasskey(pe.Encryptor.Encrypt(string(id) + string(passkey)))
}
