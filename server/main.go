package main

import (
	// "fmt"
	"log"
	"os"

	"net/http"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
) 


func NewServer() *Server {
	var (
		SUPABASE_ANON_KEY = os.Getenv("SUPABASE_ANON_KEY")
		SUPABASE_API_URL = os.Getenv("SUPABASE_API_URL")
		SUPABASE_JWT_SECRET = os.Getenv("SUPABASE_JWT_SECRET")
		SUPABASE_SERVICE_KEY = os.Getenv("SUPABASE_SERVICE_KEY") 
	) 

	if SUPABASE_ANON_KEY == "" || SUPABASE_API_URL == "" || SUPABASE_JWT_SECRET == "" || SUPABASE_SERVICE_KEY == "" {
		log.Fatal("could not initialize server")
	} 

	client, err := supabase.NewClient(SUPABASE_API_URL, SUPABASE_SERVICE_KEY, nil); 

	if err != nil {
		log.Fatalf("could not instantiate supabase client due to %+v\n", err); 
	} 

	return &Server{
		AuthClient: client,
		AnonKey: SUPABASE_ANON_KEY, 
		ApiUrl: SUPABASE_API_URL,
		JwtSecret: SUPABASE_JWT_SECRET,
		ServiceKey: SUPABASE_SERVICE_KEY, 
	} 
} 


func init() {
	err := godotenv.Load(); 
	log.Println("Init function ran sucessfully"); 
	if err != nil {
		log.Fatal("could not load env variables")
	} 
} 

func main() {
	server := NewServer(); 
	mux := http.NewServeMux(); 

	server.Middleware = [] MiddlewareLayer {
		server.authMiddleware, 
		server.allowOriginMiddleware, 
	} 
	
	// adding api endpoints 
	mux.HandleFunc("/api/createRoom", server.addMiddleware(server.createRoom)); 
	mux.HandleFunc("/api/getMyRooms", server.addMiddleware(server.getMyRooms));  
	mux.HandleFunc("/api/connectRoom", server.addMiddleware(server.connectRoom))

	log.Println("Listening to server at port http://localhost:8000 ....")
	err := http.ListenAndServe(":8000", mux); 
	if err != nil {
		log.Println("could not initiate server")
	} 
} 
