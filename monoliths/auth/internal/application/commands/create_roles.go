package commands

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type (
	CreateRolesCommand struct {
		RoleName    string
		Permissions []string
	}

	CreateRolesCommandHandler struct {
		permissionStore domain.PermissionRepository
		roleStore       domain.RoleRepository
	}
)

func NewCreateRolesCommandHandler(permissionStore domain.PermissionRepository, roleStore domain.RoleRepository) *CreateRolesCommandHandler {
	return &CreateRolesCommandHandler{
		permissionStore: permissionStore,
		roleStore:       roleStore,
	}
}

func (h *CreateRolesCommandHandler) CreateRole(ctx context.Context, command CreateRolesCommand) error {
	panic("implement me")
}
