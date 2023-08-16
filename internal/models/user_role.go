package models

type UserRole struct {
	UserName string `json:"userName"`
	RoleName string `json:"roleName"`
}

type UserRoleIdentity struct {
	UserName string `json:"userName"`
	RoleName string `json:"roleName"`
}
