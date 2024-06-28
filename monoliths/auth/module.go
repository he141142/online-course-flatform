package auth

import (
	"context"

	"drake.elearn-platform.ru/internal/systems"
	"drake.elearn-platform.ru/monoliths/auth/internal/application"
	"drake.elearn-platform.ru/monoliths/auth/internal/authpb"
	"drake.elearn-platform.ru/monoliths/auth/internal/storage/mem_store"
)

type Module struct {
}

func (m *Module) StartUp(ctx context.Context, mono systems.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, mono systems.Service) error {
	//di part
	roleRepo := mem_store.NewRoleRepository()
	permissionRepo := mem_store.NewPermissionRepository()
	server := authpb.NewMockAuthenticationServer(application.New(permissionRepo, roleRepo))
	server.RegisterRestApiEntry(ctx, *mono.HttpClient())

	return nil
}

func NewAuthModule() systems.Module {
	return &Module{}
}
