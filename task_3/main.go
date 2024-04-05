package main

import (
	"fmt"
	"keraton-task-3/database"
	"keraton-task-3/routers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai`, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("ERROR connecting database")
	} else {
		fmt.Println("Connecting and migrating database SUCCESS")
	}

	database.DbMigrate(db)

	app := fiber.New()

	//routing
	routers.Routing(app)

	app.Listen(":3000")
}
