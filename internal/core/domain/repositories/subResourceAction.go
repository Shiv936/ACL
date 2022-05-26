package repositories

import "time"

type SubResourceAction struct{
	Id string
	Sub_Resource_Id string
	Name string
	Description string
	CreatedAt time.Time
	ModifiedAt time.Time
}