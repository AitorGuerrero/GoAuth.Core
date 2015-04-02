package mongoDB

import(
	mgo "gopkg.in/mgo.v2"
)

type MongoDBConfig struct {
	Server string
}

type Connection struct {
	s *mgo.Session
	db *mgo.Database
}

func Connect(c MongoDBConfig) (*Connection, error) {
	s, err := mgo.Dial(c.Server)
	if(err != nil) {
		panic(err)
	}
	db := s.DB("badass_city_user_manager")
	return &Connection{s, db}, nil
}

func (c Connection) DB() *mgo.Database {
	return c.db
}
