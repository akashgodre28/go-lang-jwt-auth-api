package service

import (
	"UserAuth/dto"
	"UserAuth/entities"
	"UserAuth/repository"
	"UserAuth/utils"
	"errors"
	"time"
)

type UserService struct {
}

var userRepository = repository.UserRepository{}

func (UserService) CheckUserByUserName(username string) bool {
	user := userRepository.GetUserByUserName(username)
	if user.IsEmpty() {
		return false
	}
	return user.IsVerified && !user.IsDeleted && !user.IsBlocked

}

func (UserService) registerUser(request *entities.User) *entities.User {
	user := userRepository.GetUserByUserName(request.UserName)
	if user.IsEmpty() {
		hash, err := utils.HashPassword(request.Password)
		utils.CheckNilErr(err)
		request.Password = hash
		request.IsVerified = true
		request.IsDeleted = false
		request.IsBlocked = false
		request.CreatedDate = time.Now()
		request.ID = userRepository.CreateUser(request)
	}
	return request
}

func (u UserService) AuthenticateUser(request *entities.User) (*dto.Token, error) {
	user := userRepository.GetUserByUserName(request.UserName)
	if user.IsEmpty() {
		request = u.registerUser(request)
		token, err := utils.GenerateJWT(request.UserName)
		utils.CheckNilErr(err)
		return token, nil
	} else {
		match := utils.CheckPasswordHash(request.Password, user.Password)
		if match {
			token, err := utils.GenerateJWT(request.UserName)
			utils.CheckNilErr(err)
			return token, nil
		} else {
			return new(dto.Token), errors.New("invalid password")
		}
	}
}

func (u UserService) GetAllUsers() []entities.User {
	return userRepository.GetAllUsers()
}
