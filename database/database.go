package database

import (
	"fmt"
	"training-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Create() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:l1ackme@/trainingdb?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connect database successful")
	}

	// Creating the table
	if !db.HasTable(&model.Post{}) {
		db.CreateTable(&model.Post{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Post{})
	}

	// Migrate the schema
	db.AutoMigrate(&model.Post{})

	return db
}
