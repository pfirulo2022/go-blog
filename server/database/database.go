package database

import (
	"log"

	"github.com/pfirulo2022/go-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {

	dsn := "root:Agrale@tcp(127.0.0.1:3306)/fiber_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}
	log.Println("Conectado correctamente")

	if err := db.AutoMigrate(new(model.Blog)); err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	DBConn = db

}
