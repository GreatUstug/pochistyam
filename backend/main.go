package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"pochistyam/internal/app"
	"pochistyam/internal/auth"
	"pochistyam/internal/db"
	"pochistyam/internal/rooms"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	database := db.NewDB(dsn)
	defer database.Pool.Close()
	application := app.NewApp(database)
	mux := http.NewServeMux()

	dbFromPool := stdlib.OpenDBFromPool(database.Pool)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbFromPool)

	const migrationsDir = "./migrations"

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		log.Fatalf("Migrations directory not found: %s", migrationsDir)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	log.Println("Running migrations...")
	if err := goose.Up(dbFromPool, migrationsDir); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations applied successfully!")
	mux.HandleFunc("POST /createRoom", rooms.CreateRoom(application))
	mux.HandleFunc("GET /rooms", rooms.GetAllRooms(application))
	mux.HandleFunc("POST /reister", auth.Reg(application))

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
