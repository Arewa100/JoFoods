package repository

import (
	"JoFoods/src/data/models"
	"errors"
	"fmt"
	"testing"
)

func TestCreateUserRepository(test *testing.T) {
	newUserRepository := CreateUserRepository()
	if newUserRepository.MapOfUsers == nil {
		test.Error(errors.New("newUserRepository.MapOfUsers is nil, error creating user repo"))
	}
}

func TestToAddAUserToTheRepo(test *testing.T) {
	newUserRepository := CreateUserRepository()
	if newUserRepository.MapOfUsers == nil {
		test.Error(errors.New("newUserRepository.MapOfUsers is nil, error creating user repo"))
	}
	newUser := models.User{
		Username: "joseph",
	}
	newUserRepository.MapOfUsers[1] = &newUser

	if newUserRepository.MapOfUsers[1] == nil {
		test.Error(errors.New("failed to add user"))
	}
	if newUserRepository.MapOfUsers[1].Username != "joseph" {
		test.Error(errors.New("failed to add user"))
	}
}

func TestToAddTwoUsersAndFindUsersByUserName(test *testing.T) {
	newUserRepository := CreateUserRepository()
	if newUserRepository.MapOfUsers == nil {
		test.Error(errors.New("error creating repository"))
	}

	newUser := models.User{
		Username: "joseph",
	}

	newUserRepository.MapOfUsers[1] = &newUser
	foundUser, err := newUserRepository.FindUserByUserName("joseph")
	fmt.Println(foundUser)
	if foundUser == nil {
		test.Error(errors.New("user 1 not found"))
	}
	if err != nil {
		test.Error(errors.New("error finding first user"))
	}

	secondUser := models.User{
		Username: "Miracle",
	}
	newUserRepository.MapOfUsers[2] = &secondUser

	secondFoundUser, secondUserError := newUserRepository.FindUserByUserName("Miracle")
	if secondFoundUser == nil {
		test.Error(errors.New("error finding second user"))
	}
	if secondUserError != nil {
		test.Error(secondUserError)
	}

}
