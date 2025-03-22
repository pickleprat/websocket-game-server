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
	Middleware		[] MiddlewareLayer 
}


type Room struct {
	Name 		string 		`json:"name"`
	Genre 		string 		`json:"genre"`
	Description string 		`json:"description"`
	Members 	[]Member	`jsonb:"members"` 
	Owner 		string 		`json:"owner"`
} 

type Member struct {
	MemberId 	string `json:"id"`
	FullName 	string `json:"full_name"`
} 