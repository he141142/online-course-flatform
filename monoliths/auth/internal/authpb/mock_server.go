package authpb

import (
	"context"
	"net/http"

	webservers "drake.elearn-platform.ru/internal/web_servers"
	"drake.elearn-platform.ru/monoliths/auth/internal/application"
	"drake.elearn-platform.ru/monoliths/auth/internal/application/commands"
	"drake.elearn-platform.ru/monoliths/auth/internal/authpb/models"
	"drake.elearn-platform.ru/pkg/utils"
	"drake.elearn-platform.ru/pkg/utils/drk_http"
	"github.com/go-chi/chi/v5"
)

type MockAuthenticationServer struct {
	app application.App
}

func NewMockAuthenticationServer(app application.App) *MockAuthenticationServer {
	return &MockAuthenticationServer{app: app}
}

func (sv *MockAuthenticationServer) RegisterServer(ctx context.Context, httpInstance webservers.HttpChiInstance) {
	httpInstance.Route("/role", func(r chi.Router) {
		r.Post("", sv.CreateRole)
		r.Get("", sv.ListRole)
	})
}

func (sv *MockAuthenticationServer) CreateRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload models.CreateRoleRequest
	if err := drk_http.DecodeJSONBody(r, &payload); err != nil {
		drk_http.NewError(w, http.StatusBadRequest, err)
		return
	}
	permission := make([]string, 0)
	for _, p := range payload.Permissions {
		permission = append(permission, p.Name)
	}
	err := sv.app.CreateRole(ctx, commands.CreateRolesCommand{
		RoleName:                          payload.Name,
		Permissions:                       permission,
		AllowCreatePermissionWithoutExist: payload.CreatePermissionIfNotExist,
	})
	if err != nil {
		drk_http.NewError(w, http.StatusInternalServerError, err)
		return
	}
	drk_http.NewResponseWithStatus(w, http.StatusOK, nil)
}

func (sv *MockAuthenticationServer) ListRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page := utils.GetIntValue(r.URL.Query().Get("page"), 1)
	perPage := utils.GetIntValue(r.URL.Query().Get("per_page"), 5)
	roles, err := sv.app.ListRoles(ctx, page, perPage)
	if err != nil {
		drk_http.NewError(w, http.StatusInternalServerError, err)
		return
	}
	drk_http.NewResponseWithStatus(w, http.StatusOK, roles)

}
