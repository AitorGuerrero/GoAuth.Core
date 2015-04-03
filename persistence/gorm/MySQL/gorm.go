package MySQL

import (
	"github.com/AitorGuerrero/User/persistence"
	"github.com/AitorGuerrero/User/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	db gorm.DB
}

func Connect(c config.SqlDbConfig) (persistence.Connection, error) {
	db, _ := gorm.Open("mysql", c.UserName + ":" + c.Password + "@/" + c.Name + "?charset=utf8&parseTime=True&loc=Local")
	db.DB()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return &Connection{db}, nil
}
