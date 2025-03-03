package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	ws "github.com/coder/websocket"
) 


var apiLog = "%s %s Status: %s\n"
var rooms = make(Rooms, 0); 


func main() {
	rooms["chat"] = &Room{
		Id: RoomID("room-fari"), 
		Name: "fari's room", 
		Members: make([] Member, 0), 
	}

	mux := http.NewServeMux(); 
	// api handlers
	mux.HandleFunc("GET /api/ws", websockitToMe); 
	mux.HandleFunc("GET /api/getRooms", getAllRooms); 
	mux.HandleFunc("POST /api/createRoom", createRoom); 
	

	log.Println("Starting server at port 8000...");  
	err := http.ListenAndServe(":8000", mux); 
	if err != nil {
		panic(err) 
	} 
} 

func createRoom(w http.ResponseWriter, r *http.Request) {
	client, err := NewMongoClient()
	if err != nil {
		log.Printf(apiLog, "POST", "/api/createRoom", http.StatusInternalServerError); 
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError); 
	} 

	coll := client.Database(MONGO_DBNAME).Collection("rooms"); 
	result, err := coll.InsertOne(context.TODO(), nil)
	if err != nil {
		log.Printf(apiLog, "POST", "/api/createRoom", http.StatusInternalServerError); 
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} 
} 


func websockitToMe(w http.ResponseWriter, r *http.Request) {
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


