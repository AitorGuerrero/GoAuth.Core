package mgo

import(
	"github.com/AitorGuerrero/User/persistence/mongoDB"
	"github.com/AitorGuerrero/User/user"
	"github.com/AitorGuerrero/User/user/persistence"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userRepo struct {
	c *mgo.Collection
}

func Get(mc *mongoDB.Connection) userRepo {
	aRepo := userRepo {
		c: mc.DB().C("users"),
	}

	return aRepo
}

func (r userRepo) Exists(id user.Id) bool {
	var count int
	count, _ = r.c.Find(bson.M{"id": string(id)}).Count()
	return count > 0
}

func (r userRepo) Persist (u user.User) {
	r.c.Insert(u.(persistence.Serializable).Serialize())
}

func (r userRepo) Find(id user.Id) (user.User, error) {
	aUser := &user.SerializedUser{}
	r.c.Find(bson.M{"id": string(id)}).One(aUser)
	return aUser.Unserialize(), nil
}
