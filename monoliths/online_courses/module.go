package onlinecourses

import (
	"context"
	"fmt"

	"drake.elearn-platform.ru/internal/systems"
)

type Module struct {
}

func (m *Module) StartUp(ctx context.Context, mono systems.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, mono systems.Service) error {
	fmt.Println("Online courses module started")
	return nil
}

func NewOnlineCoursesModule() systems.Module {
	return &Module{}
}
