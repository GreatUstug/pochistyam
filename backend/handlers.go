package main

import (
	"encoding/json"
	"net/http"
)

type RoomResponse struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	Dirtiness     float64 `json:"dirtiness"`
	NeedsCleaning bool    `json:"needsCleaning"`
}

type CreateRoomRequest struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (app *App) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Width < 0 || req.Height < 0 {
		http.Error(w, "Name and Width and Height must be positive", http.StatusBadRequest)
		return
	}
	var room RoomResponse
	err := app.DB.Pool.QueryRow(r.Context(), `INSERT INTO rooms (user_id, name, width, height, dirtiness, needs_cleaning)
	VALUES($1, $2, $3, $4, 0.0, false)
	RETURNING id, name, width, height, dirtiness, needs_cleaning`,
		"1", req.Name, req.Width, req.Height).Scan(
		&room.ID,
		&room.Name,
		&room.Width,
		&room.Height,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(room); err != nil {
		return
	}
}
