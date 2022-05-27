package repositories

import "time"

type AccessGroup struct {
	Id          string 
	Name        string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
}
