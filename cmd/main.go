package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"user_api/config"
	"user_api/handler"
	apiError "user_api/lib/error"
	"user_api/lib/validator"
	"user_api/repositories"
	jsonobject "user_api/repositories/jsonObject"
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
	var (
		repo *repositories.Repository
	)
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

	log.Printf("initializing %s storage...", cfg.Storage)

	// init repository
	switch cfg.Storage {
	case "postgres":
		// Connect to database
		db, err := pg.NewPostgresDB()
		if err != nil {
			log.Fatalf("failed to initialize db: %s", err.Error())
		}

		// Init db repository
		repo = repositories.NewRepositoryDB(db)
		log.Printf("initialized database repository")

	case "jsonObj":
		// open json file
		users, err := jsonobject.JsonFile()
		if err != nil {
			log.Fatalf("failed to open json file: %s", err.Error())
		}

		// Init jsonObj repository
		repo = repositories.NewRepositoryJson(users)
		log.Printf("initialized jsonObj repository")

	default:
		err := fmt.Sprintf("incorrect storage: %s", cfg.Storage)
		fmt.Println(err)
		return errors.New(err)
	}

	// Init service
	userService := service.NewService(repo)

	// Init handler
	userHandler := handler.NewUsers(ctx, userService)

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

	// Create a new user.
	// Method - POST
	// Parameter content type application/json
	// request json: {"name": "string"}
	// successful response json: {"id": "integer", name: "string", "created_at": "string"}
	userRoutes.POST("/", userHandler.Create)

	// Get a user with id.
	// Method - GET
	// Parameter content type application/json
	// successful response json: {"id": "integer", name: "string", "created_at": "string"}
	userRoutes.GET("/:id", userHandler.GetUser)

	// Get all users.
	// Method - GET
	// Parameter content type application/json
	// successful response json: [{"id": "integer", name: "string", "created_at": "string"}, {},...]
	userRoutes.GET("/users", userHandler.GetAllUsers)

	// Delete a user with id.
	// Method - DELETE
	// Parameter content type application/json
	// successful response json: { "code": 200, "name": "OK", "message": "OK"}
	userRoutes.DELETE("/:id", userHandler.DeleteUser)

	// Update a user with id.
	// Method - PUT
	// Parameter content type application/json
	// request json: {"name": "string", id: "integer"}
	// successful response json: {"id": "integer", name: "string", "created_at": "string"}
	userRoutes.PUT("/:id", userHandler.Update)

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
