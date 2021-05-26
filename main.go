package main

import (
	"log"
	"os"
	"user_api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
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

	// Initialize Echo instance
	e := echo.New()

	return nil
}
