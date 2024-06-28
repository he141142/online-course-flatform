package domain

import "context"

type UserRepository interface {
	FindByID(ctx context.Context, id int) (*User, error)
	FinAll(ctx context.Context, page int, perPage int) ([]User, error)
}
