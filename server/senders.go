package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"pickleprat/ws-game-server/managers"

	ws "github.com/coder/websocket"
) 

type Sender struct {
	ctx 		context.Context
	sender 		*Member
	room 		*SupabaseRoomsResponse
	conn   		*ws.Conn 
	cm 			*managers.ConnectionManager
}

func (sender *Sender) sendMessageToRoom(msg Message) map[string] error {
	sendErrors := make(map[string] error); 
	for _, member := range sender.room.Members {
		if sender.sender.MemberId != member.MemberId {
			memberConn := sender.cm.GetConnection(member.MemberId); 
			if memberConn == nil {
				sendErrors[member.MemberId] = errors.New("member isn't present in the room") 
				continue
			} 

			buf, err := json.Marshal(msg)
			if err != nil {
				sendErrors[member.MemberId] = err 
				continue
			}  

			encodedMsg := base64.StdEncoding.EncodeToString(buf); 
			memberConn.Write(sender.ctx, ws.MessageText, []byte(encodedMsg))
		} 
	} 

	return sendErrors
} 

func NewSender(
		ctx context.Context, 
		sender *Member, 
		room *SupabaseRoomsResponse, 
		conn *ws.Conn, 
		cm *managers.ConnectionManager, 
	) *Sender {

	return &Sender{
		sender	: sender, 
		room	: room, 
		conn 	: conn, 
		cm 		: cm, 
		ctx 	: ctx, 
	}
} 