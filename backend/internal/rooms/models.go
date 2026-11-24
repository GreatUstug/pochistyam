package rooms

type RoomResponse struct {
	ID            string  `json:"id"`
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
