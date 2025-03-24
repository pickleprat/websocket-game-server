package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

)

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

	roomCreateRes := [] SupabaseRoomCreationResponse{} 
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

	rooms := make([] Room, 0); 
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

	myRooms := make([] Room, 0); 

	for _, room := range rooms {
	} 

} 
