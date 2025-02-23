package main 

import (
	"bufio" 
	"log"
	"os"
	"fmt" 
	"context" 
	"net/http"
	"encoding/json" 
	"strings" 
	"sync" 
	b64 "encoding/base64"

	ws "github.com/coder/websocket"
) 

type MemberMessage struct {
	Username string `json:"username"`
	RoomId string `json:"roomId"` 
	Msg string `json:"msg"`
} 

func main() {
	var wg sync.WaitGroup; 
	reader := bufio.NewReader(os.Stdin); 
	fmt.Printf("Username: "); 
	username, err := reader.ReadString('\n'); 
	if err != nil {
		panic(err) 
	} 

	fmt.Printf("Room ID: "); 
	roomId, err := reader.ReadString('\n'); 
	if err != nil {
		panic(err) 
	} 

	ctx := context.Background(); 
	headers := http.Header{}; 

	username = strings.TrimSpace(username); 
	roomId = strings.TrimSpace(roomId); 

	headers.Add("Username", username); 
	headers.Add("RoomId", roomId); 

	conn, _,  err := ws.Dial(ctx, 
		"ws://localhost:8000/ws", &ws.DialOptions{
			HTTPHeader : headers,   
		}, 
	); 

	if err != nil {
		panic(err); 
	} 

	defer func() {
		log.Printf("error occured while connecting!"); 
		conn.Close(ws.StatusNormalClosure, "Connection closed by the user"); 
		return 
	}() 
	
	wg.Add(2); 
	go func(ctx context.Context, username string, roomId string, conn *ws.Conn) {
		handleIncoming(ctx, username, roomId, conn); 
		wg.Done(); 
	}(ctx, username, roomId, conn)  

	go func(ctx context.Context, username string, roomId string, conn *ws.Conn) {
		handleOutgoing(ctx, username, roomId, conn); 
		wg.Done(); 
	} (ctx, username, roomId, conn)

	wg.Wait() 

} 

func handleOutgoing(ctx context.Context, username, roomId string, conn *ws.Conn) {
	scanner := bufio.NewScanner(os.Stdin); 
	for scanner.Scan() {
		msg := scanner.Text();  
		cMsg := MemberMessage{
			Username: username, 
			Msg : msg, 
			RoomId: roomId, 
		};  

		jsBuf, err := json.Marshal(cMsg); 
		if err != nil {
			log.Printf("error occured parsing json\n"); 
			continue
		} 

		encodedMsg := b64.StdEncoding.EncodeToString(jsBuf); 


		err = conn.Write(ctx, ws.MessageText, [] byte(encodedMsg)); 
		if err != nil {
			log.Printf("failed to write to server\n"); 
			continue 
		} 
	} 
} 

func handleIncoming(ctx context.Context, username, roomId string, conn *ws.Conn) {
	for {
		msgType, msg, err := conn.Read(ctx); 
		if ws.CloseStatus(err) == ws.StatusNormalClosure {
			log.Printf("connection closed normally\n"); 
			return 
		} 

		if err != nil {
			log.Println("error occured while reading msg"); 
			continue 
		} 

		if msgType != ws.MessageText {
			log.Println("message should be type text"); 
			continue 
		} 

		decodedMsg, err := b64.StdEncoding.DecodeString(string(msg)); 
		if err != nil {
			log.Printf("Error decoding from base64"); 
			continue 
		} 

		memberMsg := MemberMessage{} 
		err = json.Unmarshal(decodedMsg, &memberMsg);  
		if err != nil {
			log.Printf("error occured while parsing json buf"); 
			continue 
		} 

		fmt.Printf("%s : %s\n", memberMsg.Username, memberMsg.Msg); 
	} 
} 

