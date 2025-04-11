package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/config"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/controller"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/entity"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/repository"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/internal/usecase"
	"github.com/bakhytzhanjzz/ecommerce-platform/inventory-service/pkg/logger"
)

func runDBMigration(migrationURL string, dbSource string) error {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		return err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize logger
	log := logger.NewLogger()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Run DB migration first (before GORM)
	migrationURL := "file://migrations"
	if err := runDBMigration(migrationURL, cfg.DatabaseURL); err != nil {
		log.Fatalf("cannot run db migration: %v", err)
	}

	// Initialize database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate models (GORM auto-migration)
	if err := db.AutoMigrate(&entity.Product{}, &entity.Category{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// Initialize use cases
	productUsecase := usecase.NewProductUsecase(productRepo, log)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, log)

	// Initialize controllers
	productController := controller.NewProductController(productUsecase)
	categoryController := controller.NewCategoryController(categoryUsecase)

	// Create Gin router
	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery())
	r.Use(logger.GinLogger(log))

	// Routes
	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", productController.CreateProduct)
			products.GET("", productController.ListProducts)
			products.GET("/:id", productController.GetProduct)
			products.PATCH("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}

		categories := v1.Group("/categories")
		{
			categories.POST("", categoryController.Create)
			categories.GET("", categoryController.List)
			categories.GET("/:id", categoryController.Get)
			categories.PATCH("/:id", categoryController.Update)
			categories.DELETE("/:id", categoryController.Delete)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	log.Infof("Starting Inventory Service on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
