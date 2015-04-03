package mgo

import(
	"github.com/AitorGuerrero/User/persistence/mongoDB"
	"github.com/AitorGuerrero/User/user"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"code.google.com/p/go-uuid/uuid"
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

func (r userRepo) FindCountById(id uuid.UUID) int {
	var count int
	count, _ = r.c.Find(bson.M{"id": string(id)}).Count()
	return count
}

func (r userRepo) Persist (u interface{}) {
	r.c.Insert(u)
}

func (r userRepo) Find(id uuid.UUID) user.SerializedUser {
	aUser := user.SerializedUser{}
	r.c.Find(bson.M{"id": string(id)}).One(aUser)
	return aUser
}
