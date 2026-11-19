package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type RoomRequest struct {
	ClientID string `json:"clientId"`
	RoomID   int    `json:"roomId"`
}

type App struct {
	DB *DB
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	database := NewDB(dsn)
	defer database.Pool.Close()
	app := &App{database}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /createRoom", app.CreateRoom)
	mux.HandleFunc("GET /rooms", app.GetAllRooms)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
