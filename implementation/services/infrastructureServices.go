package services

import (
	implementationUser "github.com/AitorGuerrero/UserGo/implementation/user"
	implementationManager "github.com/AitorGuerrero/UserGo/implementation/user/manager"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)
var userSource = implementationUser.Source{map[user.Id]user.User{}}
var managerSource = implementationManager.Source{map[user.Id]manager.Manager{}}

func UserSource() user.Source {
	return &userSource
}

func ManagerSource() manager.Source {
	return &managerSource
}
