package models

type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

type PaginationModel struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

type ListRolesResponse struct {
	Roles []Role          `json:"roles"`
	Extra PaginationModel `json:"extra"`
}

type CommonResponse struct {
	Message string `json:"message"`
}
