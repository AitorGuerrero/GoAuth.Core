package newManager

import (
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
	"fmt"
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
	var id user.Id
	if id = user.ParseId(req.Id); id == nil {
		err = IncorrectIdError{}
		return
	}
	if req.Passkey == "" {
		err = EmptyPasskeyError{}
		return
	}
	m := manager.Manager{
		com.UserFactory.Make(id, req.Passkey),
		user.Namespace(req.Namespace),
		map[string]*user.User{},
	}
	fmt.Print("Manager ", m, "\n")
	fmt.Print("User", &m.User, "\n")
	m.GenerateToken()

	err = com.Source.Add(&m)
	if _, ok := err.(manager.ExistentNamespaceError); ok {
		err = ExistingNamespaceError{};
	}
	if _, ok := err.(manager.DuplicatedIdError); ok {
		err = DuplicatedIdError{};
	}

	return err;
}
