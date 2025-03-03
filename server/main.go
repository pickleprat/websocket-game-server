package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	ws "github.com/coder/websocket"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
) 

type Server struct {
	client *mongo.Client
} 

func NewServer(client *mongo.Client) *Server {
	return &Server{
		client: client, 
	} 
} 

var apiLog = "%s %s Status: %s\n"
var rooms = make(Rooms, 0); 

func handleHttpError(w http.ResponseWriter, err error, statusCode int, apiUrl, httpMethod string) {
	if err != nil {
		log.Printf(apiLog, httpMethod, apiUrl, statusCode); 
		http.Error(w, 
			fmt.Sprintf(
				"%s: Error: %s\n", 
				http.StatusText(statusCode), 
				err.Error(), 
			), 
			statusCode, 
		); 
	} 
} 


func main() {
	rooms["chat"] = &Room{
		Id: RoomID("room-fari"), 
		Name: "fari's room", 
		Members: make([] Member, 0), 
	}

	client, err := NewMongoClient()
	server := NewServer(client); 

	if err != nil {
		panic(err) 
	} 

	mux := http.NewServeMux(); 
	// api handlers
	mux.HandleFunc("GET /api/ws", server.websockitToMe); 
	mux.HandleFunc("GET /api/getRooms", server.getAllRooms); 
	mux.HandleFunc("POST /api/createRoom", server.createRoom); 
	

	log.Println("Starting server at port 8000...");  
	err = http.ListenAndServe(":8000", mux); 
	if err != nil {
		panic(err) 
	} 
} 

func (server *Server) getAllRooms(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) createRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError); 
	} 

	roomColl := server.client.Database(MONGO_DBNAME).Collection("rooms"); 
	userColl := server.client.Database(MONGO_DBNAME).Collection("users"); 

	_, err := roomColl.InsertOne(context.TODO(), nil)
	handleHttpError(w, err, http.StatusInternalServerError, "/api/createRoom", http.MethodPost); 

	// read request content 
	reqBuf, err := io.ReadAll(r.Body); 
	handleHttpError(w, err, http.StatusInternalServerError, "/api/createRoom", http.MethodPost); 

	reqJson := CreateRoomRequest{}; 
	err = json.Unmarshal(reqBuf, &reqJson);  
	handleHttpError(w, err, http.StatusInternalServerError, "/api/createRoom", http.MethodPost)

	roomId := RoomID(uuid.New().String()); 
	member := &Member{}
	filter := bson.D{{Key: "userId", Value: reqJson.Owner}};  
	userColl.FindOne(context.TODO(), filter).Decode(&member); 

	room := Room {
		Id: roomId, 
		Name: reqJson.Name,  
		Genre: reqJson.Genre, 
		Owner: reqJson.Owner, 
		Description: reqJson.Description,

	}  

	roomBuf, err := json.Marshal(room); 
	handleHttpError(w, err, http.StatusInternalServerError, "/api/createRoom", http.MethodPost); 
	encodedRoom := b64.StdEncoding.EncodeToString(roomBuf); 

	w.Header().Set("Content-Type", "application/json"); 
	w.WriteHeader(http.StatusOK); 
	json.NewEncoder(w).Encode(&CreateRoomResponse{
		RoomId: string(roomId), 
		EncodedRoom: encodedRoom, 
	}); 
} 


func (server *Server) websockitToMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed); 
	} 

	conn, err := ws.Accept(w, r, &ws.AcceptOptions{
		CompressionMode: ws.CompressionNoContextTakeover, 
	}); 


	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError); 
		return 
	} 

	defer func() {
		log.Printf("hello the connection is closing for some reason"); 
		conn.Close(ws.StatusNormalClosure, "could not connect"); 
	}() 

	ctx := r.Context(); 

	// get roomId and username via headers
	username := r.Header.Get("Username"); 
	roomId := r.Header.Get("RoomId"); 

	if username == "" || roomId == "" {
		http.Error(w, "Empty username or roomId", http.StatusBadRequest); 
		return 
	} 

	// create a member and add to the room 
	member := NewMember(username, RoomID(roomId), conn); 
	room, err := rooms.GrabBy(RoomID(roomId)); 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest); 
		return 
	} 

	room.AddMember(member); 

	for {
		// read message from user 
		msgType, msg, err := member.Conn.Read(ctx); 
		if ws.CloseStatus(err) == ws.StatusNormalClosure {
			log.Printf("Connection closed by user: %s\n", member.Username); 
			return 
		} 



		if err != nil {
			log.Printf(
			"Error occured while reading message: msgtype: %+v, msg : %+v, err: %+v\n", 
			msgType, string(msg), err); 
			continue 
		}

		if msgType != ws.MessageText {
			log.Printf("Incorrect format: Message should be text\n");  
			continue 
		} 

		fmt.Printf("Length of members : %d\n", len(room.Members));  

		// distribute messages
		for _, roomMate := range room.Members {
			if roomMate.Id() != member.Id() {
				decodedMsg, err := b64.StdEncoding.DecodeString(string(msg)); 
				if err != nil {
					log.Printf("could not decode string"); 
				} 

				memberMsg := Message{} 
				err = json.Unmarshal(decodedMsg, &memberMsg); 
				if err != nil {
					log.Printf("could not decode json"); 
				} 

				fmt.Printf("Sending message from %s to %s: %+v\n", 
					member.Id(), roomMate.Id(), memberMsg); 

				err = roomMate.Conn.Write(ctx, ws.MessageText, msg); 
				if err != nil {
					log.Printf("Could not write well to the user: %s\n", 
					roomMate.Username); 
				} 
			} 
		} 
	} 
} 


