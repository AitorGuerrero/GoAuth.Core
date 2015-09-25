package services

import (
	inMemoryUserSource "github.com/AitorGuerrero/UserGo/implementation/inMemory/userSource"
	inMemorySession "github.com/AitorGuerrero/UserGo/implementation/inMemory/session"
	"github.com/AitorGuerrero/UserGo/implementation/inMemory"
	"github.com/AitorGuerrero/UserGo/user"
)

var userSource = inMemoryUserSource.UserSource{}
var managerSource = inMemory.ManagerSource{}
var tokenSource = inMemorySession.TokenSource{}

func UserSource() user.UserSource {
	return &userSource
}

func ManagerSource() user.ManagerSource {
	return &managerSource
}

func TokenSource() user.TokenSource {
	return &tokenSource
}
