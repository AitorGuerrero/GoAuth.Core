package services

import (
	implementationUser "github.com/AitorGuerrero/UserGo/implementation/user"
	implementationManager "github.com/AitorGuerrero/UserGo/implementation/user/manager"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/user/manager"
)
var userSource = implementationUser.UserSource{make(map[user.Id]user.User)}
var managerSource = implementationManager.ManagerSource{}

func UserSource() user.UserSource {
	return &userSource
}

func ManagerSource() manager.ManagerSource {
	return &managerSource
}
