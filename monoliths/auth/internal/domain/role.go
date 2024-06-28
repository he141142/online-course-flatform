package domain

type Role struct {
	ID          int
	Name        string
	Permissions []Permission
}

func (r *Role) AddPermission(p Permission) {
	r.Permissions = append(r.Permissions, p)
}

func (r *Role) AddPermissions(p []Permission) {
	r.Permissions = append(r.Permissions, p...)
}

func NewRole(id int, name string) *Role {
	return &Role{
		ID:          id,
		Name:        name,
		Permissions: make([]Permission, 0),
	}
}

type Permission struct {
	ID          int
	Description string
	FeatureCode string
}

func NewPermission(id int, featureCode string) *Permission {
	return &Permission{
		ID:          id,
		FeatureCode: featureCode,
	}
}

func (p *Permission) AssignToRole(role *Role) {
	role.AddPermission(*p)
}
