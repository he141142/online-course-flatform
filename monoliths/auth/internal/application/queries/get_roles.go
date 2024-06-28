package queries

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type ListRoleQueryHandler struct {
	repo domain.RoleRepository
}

func NewListRoleQueryHandler(repo domain.RoleRepository) *ListRoleQueryHandler {
	return &ListRoleQueryHandler{
		repo: repo,
	}
}

func (r *ListRoleQueryHandler) ListRoles(ctx context.Context, page int, perPage int) ([]domain.Role, error) {
	return r.repo.FindAll(ctx, page, perPage)
}
