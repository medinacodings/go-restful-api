package main

import (
	"github.com/aronipurwanto/go-restful-api/app"
	"github.com/aronipurwanto/go-restful-api/controller"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/aronipurwanto/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	"log"
)

func main() {

	server := fiber.New()

	// Initialize Database
	db := app.NewDB()

	// Define the MySQL connection string
	_ = "root:Phainonphysicaldestruction#89@tcp(localhost:3306)/go_restful_api?charset=utf8mb4&parseTime=True&loc=Local"

	// Run Auto Migration (Opsional, bisa dihapus jika tidak diperlukan)

	err := db.AutoMigrate(&domain.Category{})
	err = db.AutoMigrate(&domain.Product{})
	err = db.AutoMigrate(&domain.Employee{})
	helper.PanicIfError(err)

	// Initialize Validator
	validate := validator.New()

	// Initialize Repository, Service, and Controller
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Setup Routes
	app.NewRouter(server, categoryController)

	// Start Server
	log.Println("Server running on port 8080")
	err = server.Listen(":8080")
	helper.PanicIfError(err)
}
