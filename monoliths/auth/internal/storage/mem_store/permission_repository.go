package mem_store

import (
	"drake.elearn-platform.ru/monoliths/auth/internal/domain"
)

type (
	ErrPermissionNotFound string
)

var (
	errPermissionNotFound ErrPermissionNotFound = "ErrPermissionNotFound"
)

func (ErrPermissionNotFound) Error() string {
	return string(errPermissionNotFound)
}

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

func (p *PermissionRepository) FetchLastID() (int, error) {
	if len(p.permission) == 0 {
		return 0, nil
	}
	lastID := p.permission[len(p.permission)-1].ID
	return lastID, nil
}

func (p *PermissionRepository) FindAll() ([]domain.Permission, error) {
	return p.permission, nil
}

func (p *PermissionRepository) FindByID(id int) (*domain.Permission, error) {
	for _, permission := range p.permission {
		if permission.ID == id {
			return &permission, nil
		}
	}
	return nil, errPermissionNotFound
}

func (p *PermissionRepository) FindByFeatureCode(featureCode string) (*domain.Permission, error) {
	for _, permission := range p.permission {
		if permission.FeatureCode == featureCode {
			return &permission, nil
		}
	}
	return nil, errPermissionNotFound
}

func (p *PermissionRepository) FindByFeatureCodes(featureCodes []string) ([]domain.Permission, error) {
	var permissions []domain.Permission
	for _, featureCode := range featureCodes {
		for _, permission := range p.permission {
			if permission.FeatureCode == featureCode {
				permissions = append(permissions, permission)
			}
		}
	}
	return permissions, nil
}

func (p *PermissionRepository) Create(permission *domain.Permission) error {
	lastID, err := p.FetchLastID()
	if err != nil {
		return err
	}
	permission.ID = lastID + 1
	p.permission = append(p.permission, *permission)
	return nil
}

func (p *PermissionRepository) CreateBulk(permission []*domain.Permission) error {
	for _, perm := range permission {
		err := p.Create(perm)
		if err != nil {
			return err
		}
	}
	return nil
}
