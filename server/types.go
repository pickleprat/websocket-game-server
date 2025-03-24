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