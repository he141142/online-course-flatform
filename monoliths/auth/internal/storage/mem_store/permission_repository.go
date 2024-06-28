package mem_store

import "drake.elearn-platform.ru/monoliths/auth/internal/domain"

type PermissionStore struct {
	permission []domain.Permission
}
type PermissionRepository struct {
	PermissionStore
}

var _ domain.PermissionRepository = (*PermissionRepository)(nil)

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

func (p PermissionRepository) FetchLastID() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionRepository) FindAll() ([]domain.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionRepository) FindByID(id int) (*domain.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionRepository) FindByFeatureCode(featureCode string) (*domain.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionRepository) FindByFeatureCodes(featureCodes []string) ([]domain.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionRepository) Create(permission *domain.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (p PermissionRepository) CreateBulk(permission []*domain.Permission) error {
	//TODO implement me
	panic("implement me")
}
