package data

type Action struct {
	ActionID int `json:"action_id"`
	Action string `json:"action"`
}

type Permission struct {
	PermissionID        int               `json:"permission_id"`
	Name      string             `json:"name" xml:"name"`
	Path      string             `json:"path" xml:"path"`
	Actions []*Action `json:"actions"`
}

type Role struct {
	RoleID int `json:"role_id"`
	Permissions []*Permission `json:"permissions"`
}


type AdminResponse struct {
	ID uint `gorm:"primaryKey;autoIncrement;" json:"id"`
	Email string `gorm:"size:50;not null;index;unique" json:"email"`
	FullName string `gorm:"size:50" json:"full_name"`
	Phone string `gorm:"size:10" json:"phone"`
	Roles []*Role `json:"admin_roles" xml:"admin_roles"`
}