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

func NewRole(name string) *Role {
	return &Role{
		Name:        name,
		Permissions: make([]Permission, 0),
	}
}

type Permission struct {
	ID          int
	Description string
	FeatureCode string
}

func NewPermission(featureCode string) *Permission {
	return &Permission{
		FeatureCode: featureCode,
	}
}

func (p *Permission) AssignToRole(role *Role) {
	role.AddPermission(*p)
}
