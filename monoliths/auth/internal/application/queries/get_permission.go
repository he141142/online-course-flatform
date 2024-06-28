package queries

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type ListPermissionQueryHandler struct {
	permissionRepo domain.PermissionRepository
}

func NewListPermissionQueryHandler(permissionRepo domain.PermissionRepository) *ListPermissionQueryHandler {
	return &ListPermissionQueryHandler{
		permissionRepo: permissionRepo,
	}
}

func (r *ListPermissionQueryHandler) ListPermissions(ctx context.Context, page int, perPage int) ([]domain.Permission, error) {
	return r.permissionRepo.FindAll()
}
