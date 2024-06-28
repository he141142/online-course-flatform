package commands

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type (
	CreateRolesCommand struct {
		RoleName                          string
		Permissions                       []string
		AllowCreatePermissionWithoutExist bool
	}

	CreteRolesCommandOption func(command *CreateRolesCommand)

	AllowCreatePermissionWithoutExist bool

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

func (h *CreateRolesCommandHandler) CreateRole(ctx context.Context, command CreateRolesCommand, opt ...CreteRolesCommandOption) error {
	for _, o := range opt {
		o(&command)
	}
	if command.AllowCreatePermissionWithoutExist {
		err := h.createPermissionWithoutExist(ctx, command.Permissions)
		if err != nil {
			return err
		}
	}
	permissions, err := h.permissionStore.FindByFeatureCodes(command.Permissions)
	if err != nil {
		return err
	}
	lastRoleID, err := h.roleStore.FetchLastID(ctx)
	if err != nil {
		return err
	}
	role := domain.NewRole(lastRoleID+1, command.RoleName)
	role.AddPermissions(permissions)
	return h.roleStore.Create(ctx, role)
}

func (h *CreateRolesCommandHandler) createPermissionWithoutExist(ctx context.Context, permissionPayload []string) error {
	permissions, err := h.permissionStore.FindAll()
	if err != nil {
		return err
	}
	allPermissionPresenter := make(map[string]bool)
	for _, p := range permissions {
		allPermissionPresenter[p.FeatureCode] = true
	}
	if len(allPermissionPresenter) == 0 {
		return nil
	}
	permissionToCreate := make([]*domain.Permission, 0)
	latestID, err := h.permissionStore.FetchLastID()
	if err != nil {
		return err
	}

	for _, p := range permissionPayload {
		if _, ok := allPermissionPresenter[p]; !ok {
			permission := domain.NewPermission(latestID+1, p)
			permissionToCreate = append(permissionToCreate, permission)
			latestID++
		}
	}

	return h.permissionStore.CreateBulk(permissionToCreate)
}

func (opt AllowCreatePermissionWithoutExist) Apply(command *CreateRolesCommand) {
	command.AllowCreatePermissionWithoutExist = bool(opt)
}
