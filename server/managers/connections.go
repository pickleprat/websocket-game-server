package managers

import (
	"errors"
	"sync"

	ws "github.com/coder/websocket"
) 

type ConnectionManager struct {
	Users 		map[string] *ws.Conn
	mu 			sync.Mutex
} 

func (cm *ConnectionManager) AddConnection(userId string, conn *ws.Conn) {
	cm.mu.Lock(); 
	defer cm.mu.Unlock(); 
	cm.Users[userId] = conn; 
} 

func (cm *ConnectionManager) RemoveConnection(userId string) error {
	cm.mu.Lock(); 
	defer cm.mu.Unlock(); 
	for uid := range cm.Users {
		if uid == userId {
			delete(cm.Users, userId);  
			return nil
		} 
	} 
	return errors.New("user doesn't exist in the connection manager");  
} 

func (cm *ConnectionManager) GetConnection(userId string) *ws.Conn {
	if _, ok := cm.Users[userId]; ok {
		return cm.Users[userId]
	} 
	return nil 
} 

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager {
		Users: make(map[string] *ws.Conn, 0), 
	} 
} 