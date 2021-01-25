package service

import (
	"github.com/brunobandev/reppyapi/internal/domain/property"
)

type PropertyService interface {
	Save(property property.Property) (*property.Property, error)
	FindAll() (*[]property.Property, error)
	FindById(id uint64) (*property.Property, error)
	Delete(id uint64) error
	Update(property property.Property) error
}

type propertyServiceImpl struct {
}

func NewPropertyService() PropertyService {
	return propertyServiceImpl{}
}

func (propertyServiceImpl) Save(property property.Property) (*property.Property, error) {
	return nil, nil
}

func (propertyServiceImpl) FindAll() (*[]property.Property, error) {
	return nil, nil
}

func (propertyServiceImpl) FindById(id uint64) (*property.Property, error) {
	return nil, nil
}

func (propertyServiceImpl) Delete(id uint64) error {
	return nil
}

func (propertyServiceImpl) Update(property property.Property) error {
	return nil
}
