package domain

type PermissionRepository interface {
	FetchLastID() (int, error)
	FindAll() ([]Permission, error)
	FindByID(id int) (*Permission, error)
	FindByFeatureCode(featureCode string) (*Permission, error)
	FindByFeatureCodes(featureCodes []string) ([]Permission, error)
	Create(permission *Permission) error
	CreateBulk(permission []*Permission) error
}
