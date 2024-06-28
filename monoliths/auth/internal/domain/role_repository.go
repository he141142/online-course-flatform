package domain

import "context"

type RoleRepository interface {
	FindAll(ctx context.Context, page int, perPage int) ([]Role, error)
	FindByID(ctx context.Context, id int) (*Role, error)
	Create(ctx context.Context, role *Role) error
}
