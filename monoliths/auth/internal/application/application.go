package application

import (
	"context"

	"drake.elearn-platform.ru/monoliths/auth/internal/application/commands"
	"drake.elearn-platform.ru/monoliths/auth/internal/application/queries"
	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type (
	App interface {
		Query
		Command
	}
	Query interface {
		ListRoles(ctx context.Context, page int, perPage int) ([]domain.Role, error)
		ListPermissions(ctx context.Context, page int, perPage int) ([]domain.Permission, error)
	}

	Command interface {
		CreateRole(ctx context.Context, command commands.CreateRolesCommand, opt ...commands.CreteRolesCommandOption) error
		CreatePermission(ctx context.Context, command commands.CreatePermissionCommand) error
	}

	Application struct {
		appCommand
		appQuery
	}

	appCommand struct {
		commands.CreateRolesCommandHandler
		commands.CreatePermissionCommandHandler
	}
	appQuery struct {
		queries.ListPermissionQueryHandler
		queries.ListRoleQueryHandler
	}
)

var _ App = (*Application)(nil)

func New(
	permissionStore domain.PermissionRepository,
	roleStore domain.RoleRepository,
) App {
	return &Application{
		appCommand: appCommand{
			CreateRolesCommandHandler:      *commands.NewCreateRolesCommandHandler(permissionStore, roleStore),
			CreatePermissionCommandHandler: *commands.NewCreatePermissionCommandHandler(permissionStore),
		},
		appQuery: appQuery{
			ListPermissionQueryHandler: *queries.NewListPermissionQueryHandler(permissionStore),
			ListRoleQueryHandler:       *queries.NewListRoleQueryHandler(roleStore),
		},
	}
}
