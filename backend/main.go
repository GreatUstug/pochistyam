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

	_ "pochistyam/docs"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// @title           Pochistyam API
// @version         1.0
// @description     Менеджер уборки по комнатам
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in              header
// @name            Authorization
// @description     Введите "Bearer <токен>"
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

	if err := auth.InitSecret(); err != nil {
		log.Fatal(err)
	}

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
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("POST /register", auth.Reg(application))
	mux.HandleFunc("POST /auth", auth.Auth(application))
	mux.HandleFunc("POST /createRoom", auth.AuthMiddleware(rooms.CreateRoom(application)))
	mux.HandleFunc("GET /rooms", auth.AuthMiddleware(rooms.GetAllRooms(application)))
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
