package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:7F6qtB3b7au0zYITJVkv@containers-us-west-88.railway.app:6131/railway"))
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&Users{})
	DB = db
}
