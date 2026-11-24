package rooms

import (
	"encoding/json"
	"net/http"
	"pochistyam/internal/app"
)

func CreateRoom(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("userID").(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
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
	RETURNING id, name, width, height, dirtiness, needs_cleaning`, userID, req.Name, req.Width, req.Height).Scan(
			&room.ID,
			&room.Name,
			&room.Width,
			&room.Height,
			&room.Dirtiness,
			&room.NeedsCleaning,
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
}

func GetAllRooms(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("userID").(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		rows, err := app.DB.Pool.Query(r.Context(), "SELECT id, name, width, height, dirtiness, needs_cleaning FROM rooms WHERE user_id = $1 ORDER BY id ASC", userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var rooms []RoomResponse
		for rows.Next() {
			var room RoomResponse
			if err := rows.Scan(&room.ID, &room.Name, &room.Width, &room.Height, &room.Dirtiness, &room.NeedsCleaning); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			rooms = append(rooms, room)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"rooms": rooms}); err != nil {
			return
		}
	}
}
