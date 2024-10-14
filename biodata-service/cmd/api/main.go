package main

import (
	"fmt"
	"log"
	"os"

	"myapp/config"
	"myapp/database"
	"myapp/router"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger"
	email "github.com/Abhi-singh-karuna/my_Liberary/email"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Execute the root command from the config package, which sets up the CLI and configuration
	if err := config.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger := baselogger.NewBaseLogger()
	defer logger.Sync() // Ensure logs are flushed on exit

	// Connect to database
	db, err := database.ConnectDB(cfg.SQL.Write)
	if err != nil {
		logger.Fatal("Error connecting to host database: %v", err)
	}
	logger.Infof("Host Database Connected Successfully.....")


	// Initialize validator
	validate := validator.New()

	// Initialize email service
	emailService := email.SendGridEmailService(cfg.SendGridAPIKey, cfg.SendGridFromEmail, cfg.SendGridFromName, cfg.General.Email.Enable, logger)

	// Initialize PDF generator
	// pdfGenerator := pdfoperations.NewRequestPdf(logger, cfg.UniDocMeteredKey)

	// Initialize Redis cache
	redisClient, err := database.ConnectRedis(cfg, logger)
	if err != nil {
		logger.Fatal("Error connecting to redis: %v", err)
	}

	// Create router and pass logger, emailService, and pdfGenerator
	r := router.NewRouter(db, cfg, validate, logger, emailService, redisClient)
	logger.Infof("Starting server on port %d", cfg.General.Router.Port)
	r.Run(fmt.Sprintf(":%d", cfg.General.Router.Port))
}
