package main 

type CreateRoomRequest struct {
	OwnerId 			string `json:"owner-id"`
	RoomName 			string `json:"room-name"`
	RoomGenre 			string `json:"room-genre"`
	RoomDescription     string `json:"room-description"`
} 


type CreateRoomResponse struct {
	RoomUuid 	string 		`json:"roomId"`
	RoomActive 	bool 		`json:"roomStatus"`
	CreatedAt 	string 		`json:"createdAt"`
	RoomName 	string 		`json:"roomName"`
} 

type MyRoomsRequestModel struct {
	UserUid 	string 		`json:"id"`
} 


type ConnectRoomRequest struct {
	UserId 		string 		`json:"userid"`
	RoomId 		string 		`json:"roomid"`
} 