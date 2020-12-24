package services

import (
	"github.com/saravase/golang_fiber_mvc/mvc/domain"
)

var (
	UserService = userService{}
)

type userService struct{}

func (service userService) Get(id int64) (*domain.User, error) {

	user := &domain.User{
		Id: id,
	}

	user, err := user.Get()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service userService) Save(user *domain.User) error {
	err := user.Save()
	if err != nil {
		return err
	}
	return nil
}
