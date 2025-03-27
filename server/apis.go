package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pickleprat/ws-game-server/managers"

	ws "github.com/coder/websocket"
)

var manager *managers.ConnectionManager = managers.NewConnectionManager()

func (s *Server) createRoom(w http.ResponseWriter, r *http.Request) {
	buf, err := io.ReadAll(r.Body); 
	if err != nil {
		errText := "unable to read the body" 
		http.Error(w, errText, http.StatusUnprocessableEntity);  
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText)); 
		return 
	} 

	crm := CreateRoomRequest{} 
	err = json.Unmarshal(buf, &crm); 
	if err != nil {
		errText := http.StatusText(http.StatusUnprocessableEntity); 
		http.Error(w, errText, http.StatusUnprocessableEntity); 
		s.logRequest(r, http.StatusUnprocessableEntity, errors.New(errText))
		return 
	} 

	usersBuf, count, err := s.AuthClient.From("profiles").Select("*", "exact", false).Execute();  
	members := make([] Member, 0); 
	if err != nil {
		errText := "could not retrieve from profiles"; 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
	} 

	err = json.Unmarshal(usersBuf, &members); 
	if err != nil {
		errText := "could not unmarshal members"; 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
	} 

	var owner = Member{};   
	for _, member := range members {
		if member.MemberId == crm.OwnerId {
			owner = member; 
			break 
		} 
	} 

	if count == 0 {
		errText := "no such owner exists"; 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
	} 

	room := Room{
		Owner: crm.OwnerId,
		OwnerName: owner.FullName, 
		Name: crm.RoomName,
		Genre: crm.RoomGenre,
		Description: crm.RoomDescription,
		Members: [] Member{
			{
				MemberId: crm.OwnerId, 
			 	FullName: owner.FullName, 
				CreatedAt: owner.CreatedAt, 
			},  	
		},
	} 

	data, _, err := s.AuthClient.From("rooms").Insert(room, false, "", "", "exact").Execute()
	if err != nil {
		errText := err.Error(); 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 

	roomCreateRes := [] SupabaseRoomsResponse{} 
	err = json.Unmarshal(data, &roomCreateRes)

	if err != nil {
		errText := fmt.Sprintf("could not generate room response : %s\n", string(data)); 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
	} 

	if len(roomCreateRes) > 0 {
		returnedRoom := roomCreateRes[0]

		roomRes := CreateRoomResponse{
			RoomName: returnedRoom.Name,
			RoomUuid: returnedRoom.RoomUuid,
			RoomActive: returnedRoom.RoomActive,
			CreatedAt: returnedRoom.CreatedAt,
		} 

		roomResBuf, err := json.Marshal(roomRes)
		if err != nil {
			errText := fmt.Sprintf("could not convert roomResponse to something tangible %s\n", string(roomResBuf)); 
			http.Error(w, errText, http.StatusInternalServerError); 
			s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		} 

		fmt.Fprint(w, string(roomResBuf)) 
		s.logRequest(r, http.StatusOK, nil); 

	} else {
		errText := "room was not created for some reason"; 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 

}

func (s *Server) getMyRooms(w http.ResponseWriter, r *http.Request) {
	reqBuf, err := io.ReadAll(r.Body); 
	if err != nil {
		errText := "could not read body"; 
		http.Error(w, errText, http.StatusUnprocessableEntity); 
		s.logRequest(r, http.StatusUnprocessableEntity, errors.New(errText))
		return 
	} 

	userReq := MyRoomsRequestModel{} 
	err = json.Unmarshal(reqBuf, &userReq); 
	if err != nil {
		errText :=  fmt.Sprintf("unprocessable entity : %s\n", string(reqBuf)); 
		http.Error(w, errText, http.StatusUnprocessableEntity); 
		s.logRequest(r, http.StatusUnprocessableEntity, errors.New(errText))
		return 
	} 

	rooms := make([] SupabaseRoomsResponse, 0); 
	roomBuf, liveRoomCount, err := s.AuthClient.From("rooms").Select("*", "exact", false).Execute(); 
	if err != nil {
		errText :=  fmt.Sprintf("could not fetch : %s\n", string(roomBuf)); 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 

	if liveRoomCount <= 0 {
		errText :=  "no rooms in the database"; 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 

	err = json.Unmarshal(roomBuf, &rooms); 
	if err != nil {
		errText :=  fmt.Sprintf("room could not be parsed: %s\n", string(roomBuf)); 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 

	myRooms := make([] SupabaseRoomsResponse, 0); 

	for _, room := range rooms {
		if room.Owner == userReq.UserUid{
			myRooms = append(myRooms, room)
		} 
	} 

	resBuf, err := json.Marshal(myRooms); 
	if err != nil {
		errText :=  fmt.Sprintf("room could not be parsed: %s\n", string(roomBuf)); 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 

	fmt.Fprint(w, string(resBuf)); 
	s.logRequest(r, http.StatusOK, nil)
} 

func (s *Server) connectRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errText := "only get method is allowed on this api"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 

	} 

	conn, err := ws.Accept(w, r, nil); 
	if err != nil {
		errText := err.Error() 
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	defer conn.Close(ws.StatusAbnormalClosure, "connection exit"); 

	buf, err := io.ReadAll(r.Body); 
	if err != nil {
		errText := "error reading request body"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	reqJson := ConnectRoomRequest{}
	err = json.Unmarshal(buf, &reqJson); 
	if err != nil {
		errText := "error marshalling json"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	// fetching user and room data
	profilesBuf, row, err := s.AuthClient.From("profiles").Select("*", "exact", false).Execute(); 
	if err != nil {
		errText := err.Error() + "error fetching profiles"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	if row == 0 {
		errText :=  "no user exists in the db"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	profiles := make([] Member, 0); 
	err = json.Unmarshal(profilesBuf, &profiles); 
	if err != nil {
		errText := err.Error() + "all users fetched\n"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	roomsBuf, row, err := s.AuthClient.From("rooms").Select("*", "exact", false).Execute(); 
	if err != nil {
		errText := err.Error() + "error fetching rooms"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	if row == 0 {
		errText := "no room exists in the db"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	rooms := make([] SupabaseRoomsResponse, 0); 
	err = json.Unmarshal(roomsBuf, &rooms); 
	if err != nil {
		errText := err.Error() + "could not parse the db"
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	} 

	// let's find which room this user is from 
	var theRoom *SupabaseRoomsResponse; 
	found := false 
	for _, room := range rooms {
		if room.RoomUuid == reqJson.RoomId {
			theRoom = &room; 
			found = true; 
		} 
	} 

	if !found {
		errText := "room requested by the user doesn't exist"; 
		http.Error(w, errText, http.StatusMethodNotAllowed); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
		return 
	}  

	found = false; 

	// let's also find the profile 
	var theMessageSender *Member; 
	for _, profile := range theRoom.Members {
		if profile.MemberId == reqJson.UserId {
			theMessageSender = &profile; 
			found = true; 
		} 
	} 

	ctx := r.Context(); 
	theSender := NewSender(ctx, theMessageSender, theRoom, conn, manager)
	
	for {
		tp, encodedUserMsg, err := conn.Read(ctx); 
		if tp != ws.MessageText {
			errText := err.Error() + "error while reading user message\n"
			http.Error(w, errText, http.StatusMethodNotAllowed); 
			s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
			return 
		} 

		msg := Message{}
		userMsg, err := base64.StdEncoding.DecodeString(string(encodedUserMsg)); 
		if err != nil {
			errText := err.Error() + "error decoding encoded string"
			http.Error(w, errText, http.StatusMethodNotAllowed); 
			s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
			return 
		} 

		err = json.Unmarshal(userMsg, &msg); 
		if err != nil {
			errText := err.Error() + "error unmarshalling the json"
			http.Error(w, errText, http.StatusMethodNotAllowed); 
			s.logRequest(r, http.StatusInternalServerError, errors.New(errText));  
			return 
		} 

		theSender.sendMessageToRoom(msg); 
	} 
} 
