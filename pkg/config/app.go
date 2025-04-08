package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	db, err = gorm.Open("mysql", "az-black:azmil.com@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = db

	

}

func GetDB() *gorm.DB {
	return db
}