package main 

import (
	"fmt" 
) 

type RoomID string 

type Room struct {
	Id 	RoomID
	Name 	string
	Members [] Member
} 

type Rooms map[RoomID] *Room

func (rs *Rooms) GrabBy(roomId RoomID) (*Room, error) {
	if room, ok := (*rs)[roomId]; !ok {
		return nil, fmt.Errorf("no room with id : %s exists", string(roomId)); 
	} else {
		return room, nil
	} 
} 

func (r *Room) AddMember(member *Member) bool {
	for _, roomMate := range (*r).Members {
		if member.Id() == roomMate.Id() {
			return false 
		} 
	} 

	(*r).Members = append((*r).Members, *member); 
	return true 
} 
