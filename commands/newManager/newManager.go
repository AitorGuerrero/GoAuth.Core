package newManager

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Source user.ManagerSource
	Encryptor user.PasskeyEncryptor
}

type Request struct {
	Id string
	Passkey string
	Namespace string
}

func (com Command) Execute(req Request) (error) {
	uid := user.Id(req.Id)
	p := user.Passkey(req.Passkey)
	ep := com.Encryptor.Encrypt(uid, p)
	ns := user.Namespace(req.Namespace)
	u := user.New(uid, ep)
	m := user.Manager{u, ns}
	com.Source.Add(m)

	return nil;
}
