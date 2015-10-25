package user

type Encryptor struct {}

func (Encryptor) Encrypt(i string) string {
	return i
}
