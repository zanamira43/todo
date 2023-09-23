package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zanamira43/todo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file %v", err)
	}
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// DBURL := "root:roottoor@/todoApp"
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		panic("database Connection failed")
	}

	DB = db
	db.AutoMigrate(&models.TodoDate{}, &models.Todo{}, &models.User{}, &models.Role{}, &models.Permission{})

}
