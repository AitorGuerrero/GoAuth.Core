package newManager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)

type Command struct {
	Source manager.Source
	UserFactory user.Factory
}

type Request struct {
	Id string
	Passkey string
	Namespace string
}

func (com Command) Execute(req Request) (err error) {
	id, err := user.ParseId(req.Id);
	if nil != err {
		return
	}
	m := manager.Manager{
		com.UserFactory.Make(id, user.Passkey(req.Passkey)),
		user.Namespace(req.Namespace),
		map[string]user.User{},
	}
	m.GenerateToken()
	com.Source.Add(m)

	return nil;
}
