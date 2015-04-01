package kite

import (
	"github.com/AitorGuerrero/User/services/newUser"

	"github.com/koding/kite"
)

func AddService(k *kite.Kite) {
	k.HandleFunc("new", func (r *kite.Request) (interface{}, error) {
			args, _ := r.Args.Map()
			name := args["name"].MustString()
			email := args["email"].MustString()
			password := args["password"].MustString()
			return newUser.Service(name, email, password)
		}).DisableAuthentication()
}
