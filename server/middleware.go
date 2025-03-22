package main

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
) 


func (s *Server) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer next(w, r) 
		authKey := r.Header.Get("Authorization"); 

		if authKey == "" {
			http.Error(w, "authorization token not found", http.StatusUnauthorized); 
			return 
		} 

		accessToken := strings.Trim("Bearer ", authKey); 
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error){
			return []byte(s.JwtSecret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			http.Error(w, "authentication token is invalid", http.StatusUnauthorized)
			return 
		} 

		if _, ok := token.Claims.(jwt.MapClaims); !ok {
			http.Error(w, "unable to map claims", http.StatusUnauthorized); 
			return
		}  
		
	}  
}  

func (s *Server) allowOriginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*"); 	
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        
        next(w, r)
	} 
} 

func (s *Server) addMiddleware(apiFunction http.HandlerFunc) http.HandlerFunc {
	for _, mw := range s.Middleware {
		apiFunction = mw(apiFunction); 
	} 
	return apiFunction
} 
