package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:7F6qtB3b7au0zYITJVkv@tcp(railway:6131)/telegoGPT"))
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&Users{})
	DB = db
}
