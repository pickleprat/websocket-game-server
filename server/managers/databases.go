package managers

type Component interface {
	GetId() string 
} 

type DatabaseManager interface {
	Find(column string) (Component, error) 
	Select(column string) ([] Component, error)  
} 