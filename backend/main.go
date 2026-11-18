package main

import (
	"encoding/json"
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

	// заглушка
	mux.HandleFunc("POST /rooms/get", func(w http.ResponseWriter, r *http.Request) {
		var req RoomRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := RoomResponse{
			ID:            req.RoomID,
			Name:          "Комната",
			Width:         20,
			Height:        20,
			Dirtiness:     0.8,
			NeedsCleaning: true,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
