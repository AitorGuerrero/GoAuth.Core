package newManager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)

type Command struct {
	Source manager.ManagerSource
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
	ns := manager.Namespace(req.Namespace)
	m := manager.Manager{u, ns}
	com.Source.Add(m)

	return nil;
}
