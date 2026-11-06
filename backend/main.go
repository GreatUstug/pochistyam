package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RoomRequest struct {
	ClientID string `json:"clientId"`
	RoomID   int    `json:"roomId"`
}

type RoomResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Dirtiness float64 `json:"dirtiness"`
}

func main() {

	http.HandleFunc("POST /rooms/get", func(w http.ResponseWriter, r *http.Request) {
		var req RoomRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := RoomResponse{
			ID:        req.RoomID,
			Name:      "Комната",
			Width:     20,
			Height:    20,
			Dirtiness: 0.5,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
