package repositories

import "time"

type AccessGroupPermission struct {
	Access_Group_Id        string
	Resource_Id            string
	Sub_Resource_Id        string
	Sub_Resource_Action_Id string
	CreatedAt              time.Time
	ModifiedAt             time.Time
}
