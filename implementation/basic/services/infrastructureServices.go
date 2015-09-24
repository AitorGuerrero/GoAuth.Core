package services

import (
	inMemoryUserSource "github.com/AitorGuerrero/UserGo/implementation/inMemory/userSource"
	"github.com/AitorGuerrero/UserGo/user"
)

var userSource = inMemoryUserSource.UserSource{}

func UserSource() user.UserSource {
	return &userSource
}
