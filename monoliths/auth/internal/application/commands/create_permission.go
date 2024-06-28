package commands

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type CreatePermissionCommand struct {
	FeatureCode string
	Description string
}

type CreatePermissionCommandHandler struct {
	permissionRepository domain.PermissionRepository
}

func NewCreatePermissionCommandHandler(permissionRepository domain.PermissionRepository) *CreatePermissionCommandHandler {
	return &CreatePermissionCommandHandler{
		permissionRepository: permissionRepository,
	}
}

func (a *CreatePermissionCommandHandler) CreatePermission(ctx context.Context, command CreatePermissionCommand) error {
	//TODO implement me
	panic("implement me")
}
