package models

type Permission struct {
	Name string `json:"name"`
}

type CreateRoleRequest struct {
	Name                       string       `json:"name"`
	Permissions                []Permission `json:"permissions"`
	CreatePermissionIfNotExist bool         `json:"create_permission_if_not_exist"`
}
