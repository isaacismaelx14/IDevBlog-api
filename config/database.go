package config

import (
	"os"

	"github.com/isaacismaelx14/blog-api/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func Init() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(schemas.Post{})
	db.AutoMigrate(schemas.User{})
	DataBase = db
}
