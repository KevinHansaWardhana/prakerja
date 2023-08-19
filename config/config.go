package config

import (
	"prakerja/structs"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:NR1gnsGxsJ9jz1ljgghh@tcp(prakerja-production-f6dd.up.railway.app:7048)/railway?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
