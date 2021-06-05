package database

import (
	"fmt"
	"github/sing3demons/go-fiber-admin/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(db)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, name)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	database.AutoMigrate(&models.User{})
	// database.Migrator().DropTable(&models.User{})
	DB = database
}
