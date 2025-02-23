package main 

import (
	ws "github.com/coder/websocket"
) 

type Member struct {
	Username 	string 
	RoomId 		RoomID
	Conn 		*ws.Conn
} 

func (m *Member) Id() string {
	return m.Username; 
} 

func NewMember(username string, roomId RoomID, conn *ws.Conn) *Member {
	return &Member{
		Username: username, 
		RoomId: roomId, 
		Conn: conn, 
	}
} 
