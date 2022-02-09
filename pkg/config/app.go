package config

import (
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/sqlite"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	// TODO: added mysql conf -: ?charset=utf8&parseTime=True&loc=Local
	// d, err := gorm.Open("mysql", "---")
	d, err := gorm.Open("sqlite", "db.sqlite")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
