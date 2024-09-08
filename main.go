package main

import (
	"database/sql"
	"errors"
	"github.com/carolineeey/multifinance-api/api"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	globalConfig := NewConfigDefault()

	// Capture connection properties.
	cfg := mysql.Config{
		User:   globalConfig.MySqlUsername,
		Passwd: globalConfig.MySqlPassword,
		Net:    "tcp",
		Addr:   globalConfig.MySqlAddress,
		DBName: globalConfig.MySqlDbName,
	}

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	apiClient := api.NewClient(db, "0.0.0.0:8080")

	logWriter := os.Stdout

	router := createRouter(apiClient, logWriter)
	server := &http.Server{
		Addr:    globalConfig.ListenAddress,
		Handler: router,
	}

	log.Printf("Starting server on %s", globalConfig.ListenAddress)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
