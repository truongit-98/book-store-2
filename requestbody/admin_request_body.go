package requestbody

type RoleRequest struct {
	Name   string `json:"name" xml:"name"`
	Active bool   `json:"active"`
}

type RolePutRequest struct {
	RoleId uint   `json:"role_id" xml:"role_id"`
	Name   string `json:"name" xml:"name"`
	Active bool   `json:"active"`
}

type RoleUserRequest struct {
	UserId uint `json:"user_id"`
	RoleId uint `json:"role_id" xml:"role_id"`
	Active bool `json:"active" xml:"active"`
}

type role struct {
	RoleId uint `json:"role_id" xml:"role_id"`
	Active bool `json:"active" xml:"active"`
}

type MultiRoleUserRequest struct {
	UserId uint `json:"user_id"`
	Roles []role `json:"roles"`
}

type PermissionRequest struct {
	Name string `json:"name" xml:"name"`
	Path string `json:"path" xml:"path"`
	Active bool `json:"active" xml:"active"`

}
type PermissionPutRequest struct {
	PermissionId uint  `json:"permission_id" xml:"permission_id"`
	Name         string `json:"name" xml:"name"`
	Path         string `json:"path" xml:"path"`
	Active bool `json:"active" xml:"active"`

}

type RolePermissionRequest struct {
	RoleId       uint  `json:"role_id" xml:"role_id"`
	PermissionId uint  `json:"permission_id" xml:"permission_id"`
	Active       bool   `json:"active" xml:"active"`
}

type control struct {
	ControlId uint `json:"control_id" xml:"control_id"`
	Active    bool  `json:"active" xml:"active"`
}

type permission struct {
	PermissionId uint     `json:"permission_id" xml:"permission_id"`
	Controls     []control `json:"controls" xml:"controls"`
}

type MultiRolePermissionRequest struct {
	RoleId      uint        `json:"role_id" xml:"role_id"`
	Permissions []permission `json:"permissions" xml:"permissions"`
}
type RolePermissionPutRequest struct {
	Id           uint  `json:"id" xml:"id"`
	RoleId       uint  `json:"role_id" xml:"role_id"`
	PermissionId uint  `json:"permission_id" xml:"permission_id"`
	Active       bool   `json:"active" xml:"active"`
}

type AuthorControlRequest struct {
	Name         string `json:"name" xml:"name"`
	Action       string `json:"action" xml:"action"`
	Active       bool   `json:"active" xml:"active"`
}
type AuthorControlPutRequest struct {
	Id           uint  `json:"id" xml:"id"`
	Name         string `json:"name" xml:"name"`
	Action       string `json:"action" xml:"action"`
	Active       bool   `json:"active" xml:"active"`
}


type RolePermsControlRequest struct {
	RoleId uint `json:"role_id"`
	PermissionId uint `json:"permission_id"`
	Controls []control `json:"controls"`
}

type AccountRequest struct {
	Email string ` json:"email"`
	Password string ` json:"password"`
	FullName string `json:"full_name"`
	Address string ` json:"address"`
	Phone string ` json:"phone"`
}

type AdminRequest struct {
	Email string ` json:"email"`
	FullName string `json:"full_name"`
	Phone string ` json:"phone"`
}

type AdminPutRequest struct {
	ID int `json:"id"`
	Email string ` json:"email"`
	FullName string `json:"full_name"`
	Phone string ` json:"phone"`
}

type AccountLoginRequest struct {
	Email string ` json:"email"`
	Password string ` json:"password"`
}

type AccountPutRequest struct {
	Id uint `json:"id"`
	FullName string `json:"full_name"`
	Address string ` json:"address"`
	Phone string ` json:"phone"`
}


