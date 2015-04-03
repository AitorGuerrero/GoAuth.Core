package kite

import (
	"github.com/AitorGuerrero/User/services/findById"
	userRepo "github.com/AitorGuerrero/User/user/persistence"

	"github.com/koding/kite"
)

func AddService(k *kite.Kite, ur userRepo.UserRepo) {
	k.HandleFunc("find-by-id", func (r *kite.Request) (interface{}, error) {
			id := r.Args.One().MustString()
			return findById.Service(id, ur)
		}).DisableAuthentication()
}
