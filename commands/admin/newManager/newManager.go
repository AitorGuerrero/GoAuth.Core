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
	if req.Passkey == "" {
		err = EmptyPasskeyError{}
		return
	}
	if com.Source.ExistsWithNamespace(user.Namespace(req.Namespace)) {
		err = ExistingNamespaceError{}
		return
	}
	m := manager.Manager{
		com.UserFactory.Make(id, req.Passkey),
		user.Namespace(req.Namespace),
		map[string]*user.User{},
	}
	m.GenerateToken()
	err = com.Source.Add(&m)

	return err;
}
