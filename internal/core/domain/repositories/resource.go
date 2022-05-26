package repositories

import "time"

type Resource struct {
	Id          string
	Name        string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
}
