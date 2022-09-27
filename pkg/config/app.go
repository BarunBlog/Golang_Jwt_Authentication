package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_test_db?charset=utf8&parseTime=True&loc=Local")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	db = d

}

func GetDb() *gorm.DB {
	return db
}
