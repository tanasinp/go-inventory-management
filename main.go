package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tanasinp/go-inventory-management/adapters"
	"github.com/tanasinp/go-inventory-management/core"
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

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portStr, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	productRepo := adapters.NewGormProductRepository(db)
	productService := core.NewProductService(productRepo)
	productHandler := adapters.NewHttpProductHandler(productService)

	db.AutoMigrate(&database.Product{}, &database.Category{}, &database.Supplier{}, &database.ProductCategory{}, &database.User{})
	fmt.Println("Automigrate Successful")

	app := fiber.New()

	app.Post("/supplier", productHandler.CreateSupplierFiber)
	app.Post("/category", productHandler.CreateCategoryFiber)
	app.Get("/supplier", productHandler.GetAllSupplierFiber)
	app.Get("/category", productHandler.GetAllCategoryFiber)
	app.Post("/product", productHandler.CreateProductFiber)
	app.Get("/product/:id", productHandler.GetProductByIDFiber)
	app.Get("/product", productHandler.GetAllProductFiber)
	app.Get("/category/:id/product", productHandler.GetAllProductOfCategoryFiber)
	app.Get("/supplier/:id/product", productHandler.GetAllProductOfSupplierFiber)
	app.Put("/supplier/:id", productHandler.UpdateSupplierFiber)
	app.Put("/product/:id", productHandler.UpdateProductByIDFiber)
	app.Delete("/product/:id", productHandler.DeleteProductByIDFiber)

	app.Listen(":8000")
}
