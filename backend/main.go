package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

type RoomRequest struct {
	ClientID string `json:"clientId"`
	RoomID   int    `json:"roomId"`
}

type App struct {
	DB *DB
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	database := NewDB(dsn)
	defer database.Pool.Close()
	app := &App{database}
	mux := http.NewServeMux()

	db := stdlib.OpenDBFromPool(database.Pool)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	const migrationsDir = "./migrations"

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		log.Fatalf("Migrations directory not found: %s", migrationsDir)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	log.Println("Running migrations...")
	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations applied successfully!")
	mux.HandleFunc("POST /createRoom", app.CreateRoom)
	mux.HandleFunc("GET /rooms", app.GetAllRooms)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
