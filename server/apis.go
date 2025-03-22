package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
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

	room := Room{
		Owner: crm.OwnerId,
		Name: crm.RoomName,
		Genre: crm.RoomGenre,
		Description: crm.RoomDescription,
		Members: [] Member{
			{MemberId: crm.OwnerId, FullName: ""},  	
		},
	} 

	_, _, err = s.AuthClient.From("rooms").Insert(room, false, "", "", "").Execute()
	log.Printf("%s\n", err)
	if err != nil {
		errText := http.StatusText(http.StatusInternalServerError); 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 
}