package config

import (
	"github.com/jinz/gorm"
	_ "github.com/jinz/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "az-black: azmil.adhim20@gmail.com/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
