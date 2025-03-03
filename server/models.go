package main 

type CreateRoomRequest struct {
	Name 		string `json:"room-name"`
	Genre 		string `json:"genre"`
	Description string `json:"description"`
	Owner 		string `json:"owner"`
} 

type CreateRoomResponse struct {
	RoomId 			string `json:"roomId"`
	EncodedRoom 	string `json:"encodedRoom"`
} 