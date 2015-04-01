package kite

import (
	"github.com/AitorGuerrero/BadassCity/user/services/isValid"

	"github.com/koding/kite"
)

func AddService(k *kite.Kite) {
	k.HandleFunc("is-valid", func (r *kite.Request) (interface{}, error) {
			id := r.Args.One().MustString()
			return isValid.Service(id)
		}).DisableAuthentication()
}

