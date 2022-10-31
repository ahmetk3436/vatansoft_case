package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func init() {

	// Connection stringi yaratılır
	dsn := "root:12345678@tcp(127.0.0.1:3306)/vatansoft?charset=utf8mb4&parseTime=True&loc=Local"

	// Eğer Heroku üzerinde bir PostgreSQL'e sahipseniz, bu ayarlamaları yapmak yerine doğrudan
	// heroku tarafından verilen database url'i kullanabilirsiniz
	// dbUri := fmt.Sprintf("postgres://xxxxx@xxx.compute.amazonaws.com:5432/ddjkb1easq2mec") // Database url

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		fmt.Println(err.Error())
	}

	db = conn
	db.Debug().AutoMigrate(&Student{}) //Database migration
	db.Debug().AutoMigrate(&Plan{})    //Database migration
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
