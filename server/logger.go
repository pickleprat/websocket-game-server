package main

import (
	"log"
	"net/http"
)

var (
 	RED = "\033[31m" 
	GREEN = "\033[32m" 	
) 

func (s *Server) logRequest(r *http.Request, statuscode int, err error) {
	if err != nil {
		log.Printf("%s %s %s %s %s %d %s %s\n", RED, r.Method, RED, r.URL.Path, RED, statuscode, RED, err.Error());   
	} else {
		log.Printf("%s %s %s %s %s %d SUCCESS %s\n", GREEN, r.Method, GREEN, r.URL.Path, GREEN, statuscode, GREEN);   
	}  
} 