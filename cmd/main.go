package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"user_api/config"
	"user_api/repositories/pg"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	// ctx := context.Background()

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
	fmt.Sprintln(db)

	u := pg.NewUserPostgres(db)
	id, err := u.CreateUser("Petr")
	if err != nil {
		log.Printf("Create user:%v ERROR: %v\n", id, err)
	}
	fmt.Printf("id: %v\n", id)
	err = u.DeleteUser(1)
	if err != nil {
		fmt.Println(err)
	}
	us, err := u.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("USERS:", us)

	// Initialize Echo instance
	e := echo.New()

	// Init handler

	// Set middleware

	// API v1
	// v1 := e.Group("/v1")

	// set routes
	// userRoutes := v1.Group("/user")
	// fmt.Println(userRoutes)
	// userRoutes.POST("/")

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
