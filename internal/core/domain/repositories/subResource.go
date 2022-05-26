package repositories

import "time"

type SubResource struct {
	Id          string
	Resource_Id string
	Name        string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
}
