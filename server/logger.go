package main

import (
	"log"
	"net/http"
)

func (s *Server) logRequest(r *http.Request, statuscode int, err error) {
	log.Printf("%s %s %d %s\n", r.Method, r.URL.Path, statuscode, err.Error()); 
} 