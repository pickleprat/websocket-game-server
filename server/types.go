package main

import (
	"net/http"

	"github.com/supabase-community/supabase-go"
)

type MiddlewareLayer func(http.HandlerFunc) http.HandlerFunc

type Server struct {
	AuthClient 		*supabase.Client
	AnonKey 		string 
	ApiUrl 			string 
	JwtSecret 		string 
	ServiceKey 		string 
	Middleware		[] MiddlewareLayer 
}

// database components 
type Room struct {
	OwnerName 	string 		`json:"owner-name"`
	Name 		string 		`json:"name"`
	Genre 		string 		`json:"genre"`
	Description string 		`json:"description"`
	Members 	[]Member	`json:"members"` 
	Owner 		string 		`json:"owner"`
} 

type Member struct {
	MemberId 	string `json:"id"`
	FullName 	string `json:"full_name"`
	CreatedAt  	string `json:"created_at"`
} 

type Message struct {
	Msg 		string `json:"msg"`
	UserUid 	string `json:"userId"`
	RoomId 		string `json:"roomId"`
	Name 		string `json:"name"`
} 

type SupabaseRoomsResponse struct {
	Room 
	RoomUuid 	string 		`json:"id"`
	CreatedAt 	string 		`json:"created_at"`
	RoomActive	bool 		`json:"roomActive"`
} 

func (srr *SupabaseRoomsResponse) GetId() string {
	return srr.RoomUuid
} 

func (msg *Message) GetId() string {
	return msg.UserUid + msg.RoomId
} 

func (m *Member) GetId() string {
	return m.MemberId
} 

func (r * Room) GetId() string {
	return r.Owner
} 


