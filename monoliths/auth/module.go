package auth

import (
	"context"

	"drake.elearn-platform.ru/internal/systems"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono systems.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, mono systems.Service) error {
	//authpb.NewMockAuthenticationServer()
	panic("implement me")
}
