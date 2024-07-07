package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tanasinp/go-inventory-management/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	//Get from env
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")      // default PostgreSQL port
	user := os.Getenv("DB_USER")         // as defined in docker-compose.yml
	password := os.Getenv("DB_PASSWORD") // as defined in docker-compose.yml
	dbname := os.Getenv("DB_NAME")       // as defined in docker-compose.yml

	// fmt.Printf("host=%s port=%s user=%s password=%s dbname=%s\n", host, port, user, password, dbname)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting DB_PORT to int")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&database.Product{}, &database.Category{}, &database.Supplier{}, &database.ProductCategory{})
	fmt.Println("Automigrate Successful")

	app := fiber.New()

	app.Listen(":8000")
}
