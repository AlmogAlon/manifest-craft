package database

import (
	"fmt"
	"log"
	"manifest-craft/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		log.Fatal("Failed to connect to database!", error)
		os.Exit(2)
	}

	log.Println("Connected to database!")

	db.Logger = db.Logger.LogMode(logger.Info)

	log.Println("running migrations...")
	db.AutoMigrate(&models.Manifest{}, &models.Component{})

	DB = Dbinstance{Db: db}
}
