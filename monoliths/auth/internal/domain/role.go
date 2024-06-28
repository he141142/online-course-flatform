package domain

type Role struct {
	ID          int
	Name        string
	Permissions []Permission
}

type Permission struct {
	ID          int
	Description string
	FeatureCode string
}
