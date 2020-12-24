package domain

import (
	"errors"
	"log"
)

func (user *User) Save() error {

	if user == nil {
		return errors.New("Invalid user to save")
	}
	log.Printf("User: %v\n", user)

	result := dbClient.Create(user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (user *User) Get() (*User, error) {

	var _user = new(User)
	result := dbClient.Where("id = ?", user.Id).Find(&_user)
	if err := result.Error; err != nil {
		return nil, err
	}

	if _user.Id == 0 {
		return nil, errors.New("User not found")
	}

	return _user, nil
}
