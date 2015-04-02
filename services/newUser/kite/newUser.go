package kite

import (
	"github.com/AitorGuerrero/User/services/newUser"
	userRepo "github.com/AitorGuerrero/User/user/persistence"

	"github.com/koding/kite"
)

func AddService(k *kite.Kite, ur userRepo.UserRepo) {
	k.HandleFunc("new", func (r *kite.Request) (interface{}, error) {
			args, _ := r.Args.Slice()
			name := args[0].MustString()
			email := args[1].MustString()
			password := args[2].MustString()
			return newUser.Service(name, email, password, ur)
		}).DisableAuthentication()
}
