package domain

import "context"

type RoleRepository interface {
	FetchLastID(ctx context.Context) (int, error)
	FindAll(ctx context.Context, page int, perPage int) ([]Role, error)
	FindByID(ctx context.Context, id int) (*Role, error)
	Create(ctx context.Context, role *Role) error
}
