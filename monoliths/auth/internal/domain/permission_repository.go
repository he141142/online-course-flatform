package domain

type PermissionRepository interface {
	FindAll() ([]Permission, error)
	FindByID(id int) (*Permission, error)
	FindByFeatureCode(featureCode string) (*Permission, error)
	FindByFeatureCodes(featureCodes []string) ([]Permission, error)
}
