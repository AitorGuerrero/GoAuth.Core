package kite

import (
	"github.com/AitorGuerrero/User/services/isValid"
	"github.com/AitorGuerrero/User/persistence/userRepo"

	"github.com/koding/kite"
)

func AddService(k *kite.Kite, ur userRepo.UserRepo) {
	k.HandleFunc("is-valid", func (r *kite.Request) (interface{}, error) {
			id := r.Args.One().MustString()
			return isValid.Service(id, ur)
		}).DisableAuthentication()
}

