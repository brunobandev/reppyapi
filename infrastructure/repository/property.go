package repository

import "github.com/brunobandev/reppyapi/internal/domain/property"

type PropertyRepository interface {
	Save(property property.Property) (*property.Property, error)
}

type propertyRepositoryImpl struct {
}

func NewPropertyRepository() PropertyRepository {
	return propertyRepositoryImpl{}
}

func (propertyRepositoryImpl) Save(property property.Property) (*property.Property, error) {
	return nil, nil
}
