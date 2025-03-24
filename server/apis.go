package main

import (
	"encoding/json"
	"errors"
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

	_, _, err = s.AuthClient.From("rooms").Insert(room, false, "", "", "exact").Execute()
	if err != nil {
		errText := err.Error(); 
		http.Error(w, errText, http.StatusInternalServerError); 
		s.logRequest(r, http.StatusInternalServerError, errors.New(errText))
		return 
	} 

	s.logRequest(r, http.StatusOK, nil); 
}