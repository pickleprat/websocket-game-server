package main 

type CreateRoomRequest struct {
	OwnerId 			string `json:"owner-id"`
	RoomName 			string `json:"room-title"`
	RoomGenre 			string `json:"room-genre"`
	RoomDescription     string `json:"room-description"`
} 