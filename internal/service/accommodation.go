package service

import "github.com/brunobandev/reppyapi/internal/domain/accommodation"

type AccommodationService interface {
	Save(accommodation accommodation.Accommodation) (*accommodation.Accommodation, error)
	FindAll() (*[]accommodation.Accommodation, error)
	FindById(id uint64) (*accommodation.Accommodation, error)
	Delete(id uint64) error
	Update(accommodation accommodation.Accommodation) error
}

type accommodationServiceImpl struct {
}

func NewAccommodationService() AccommodationService {
	return accommodationServiceImpl{}
}

func (accommodationServiceImpl) Save(accommodation accommodation.Accommodation) (*accommodation.Accommodation, error) {
	return nil, nil
}

func (accommodationServiceImpl) FindAll() (*[]accommodation.Accommodation, error) {
	return nil, nil
}

func (accommodationServiceImpl) FindById(id uint64) (*accommodation.Accommodation, error) {
	return nil, nil
}

func (accommodationServiceImpl) Delete(id uint64) error {
	return nil
}

func (accommodationServiceImpl) Update(accommodation accommodation.Accommodation) error {
	return nil
}
