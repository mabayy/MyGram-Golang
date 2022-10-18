package database

import (
	"final-projek/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=My-Gram sslmode=disable port=5432 TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	fmt.Println("Success conecting to database")
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	return db
}
