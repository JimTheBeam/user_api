package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"user_api/config"
	"user_api/handler"
	apiError "user_api/lib/error"
	"user_api/lib/validator"
	"user_api/repositories"
	"user_api/repositories/pg"
	"user_api/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	ctx := context.Background()

	// config
	cfg := config.Get()

	// load logfile
	if cfg.LogPath != "" {
		log.Printf("Log file is: %s", cfg.LogPath)
		lf, err := os.OpenFile(cfg.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			log.Printf("Open logfile: %s", err)
			log.Fatal(err)
		}
		defer lf.Close()
		log.SetOutput(lf)
	}
	log.Println("app is starting...")

	// Connect to database
	db, err := pg.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Init repository
	repo := repositories.NewRepository(db)

	// Init service
	userService := service.NewService(ctx, repo)

	// Init handler
	UserHandler := handler.NewUsers(ctx, userService)

	// Initialize Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.HTTPErrorHandler = apiError.Error

	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API v1
	v1 := e.Group("/v1")

	// set routes
	userRoutes := v1.Group("/user")
	userRoutes.POST("/", UserHandler.Create)
	userRoutes.GET("/:id", UserHandler.GetUser)
	userRoutes.GET("/users", UserHandler.GetAllUsers)
	userRoutes.GET("/delete/:id", UserHandler.DeleteUser)
	userRoutes.PUT("/:id", UserHandler.Update)

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
