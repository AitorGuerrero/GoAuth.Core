package manager

import "github.com/AitorGuerrero/UserGo/user"

type PlainManager struct {
	User user.PlainUser
	Namespace user.Namespace
}

func (m Manager) Plain() PlainManager {
	return PlainManager {
		User: m.User.Plain(),
		Namespace: string(m.Namespace),
	}
}
