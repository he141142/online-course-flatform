package mem_store

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type (
	ErrRoleNotFound struct{}
)

type RoleStore struct {
	roles []domain.Role
}

type RoleRepository struct {
	RoleStore
}

var _ domain.RoleRepository = (*RoleRepository)(nil)

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{RoleStore{make([]domain.Role, 0)}}
}

func (r *RoleRepository) FetchLastID(ctx context.Context) (int, error) {
	if len(r.roles) == 0 {
		return 0, nil
	}
	lastID := r.roles[len(r.roles)-1].ID
	return lastID, nil
}

func (r *RoleRepository) FindAll(ctx context.Context, page int, perPage int) ([]domain.Role, error) {
	return r.roles, nil
}

func (r *RoleRepository) FindByID(ctx context.Context, id int) (*domain.Role, error) {
	for _, role := range r.roles {
		if role.ID == id {
			return &role, nil
		}
	}
	return nil, ErrRoleNotFound{}
}

func (r *RoleRepository) Create(ctx context.Context, role *domain.Role) error {
	lastID, err := r.FetchLastID(ctx)
	if err != nil {
		return err
	}
	role.ID = lastID + 1
	r.roles = append(r.roles, *role)
	return nil
}

func (ErrRoleNotFound) Error() string {
	return "role not found"
}
