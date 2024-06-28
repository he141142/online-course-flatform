package mem_store

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type RoleStore struct {
	roles []domain.Role
}

type RoleRepository struct {
	RoleStore
}

var _ domain.RoleRepository = (*RoleRepository)(nil)

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (r RoleRepository) FetchLastID(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleRepository) FindAll(ctx context.Context, page int, perPage int) ([]domain.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleRepository) FindByID(ctx context.Context, id int) (*domain.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleRepository) Create(ctx context.Context, role *domain.Role) error {
	//TODO implement me
	panic("implement me")
}
