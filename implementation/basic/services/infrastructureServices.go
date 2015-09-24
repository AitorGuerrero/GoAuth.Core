package services

import (
	inMemoryUserSource "github.com/AitorGuerrero/UserGo/implementation/inMemory/userSource"
	inMemorySession "github.com/AitorGuerrero/UserGo/implementation/inMemory/session"
	"github.com/AitorGuerrero/UserGo/user"
	"github.com/AitorGuerrero/UserGo/session"
)

var userSource = inMemoryUserSource.UserSource{}
var tokenSource = inMemorySession.TokenSource{}

func UserSource() user.UserSource {
	return &userSource
}

func TokenSource() session.TokenSource {
	return &tokenSource
}
