package service

import "github.com/brunobandev/reppyapi/internal/domain/user"

type UserService interface {
	Save(user user.User) (*user.User, error)
	FindAll() (*[]user.User, error)
	FindById(id uint64) (*user.User, error)
	Delete(id uint64) error
	Update(user user.User) error
}

type userServiceImpl struct {
}

func NewUserService() UserService {
	return userServiceImpl{}
}

func (userServiceImpl) Save(user user.User) (*user.User, error) {
	return nil, nil
}

func (userServiceImpl) FindAll() (*[]user.User, error) {
	return nil, nil
}

func (userServiceImpl) FindById(id uint64) (*user.User, error) {
	return nil, nil
}

func (userServiceImpl) Delete(id uint64) error {
	return nil
}

func (userServiceImpl) Update(user user.User) error {
	return nil
}
