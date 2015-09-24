package newManager

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type Command struct {
	Source user.ManagerSource
	Factory user.Factory
}

type Request struct {
	Id string
	Passkey string
	Namespace string
}

func (com Command) Execute(req Request) (error) {
	uid := user.Id(req.Id)
	p := user.Passkey(req.Passkey)
	u := com.Factory.Make(uid, p)
	ns := user.Namespace(req.Namespace)
	m := user.Manager{u, ns}
	com.Source.Add(m)

	return nil;
}
