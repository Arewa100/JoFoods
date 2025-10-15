package repository

import (
	"JoFoods/src/data/models"
	"errors"
)

type UserRepository struct {
	MapOfUsers map[int]*models.User
}

func CreateUserRepository() *UserRepository {
	newUserRepository := new(UserRepository)
	newUserRepository.MapOfUsers = make(map[int]*models.User)
	return newUserRepository
}

func (userRepository *UserRepository) FindUserByUserName(username string) (*models.User, error) {
	for _, user := range userRepository.MapOfUsers {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
