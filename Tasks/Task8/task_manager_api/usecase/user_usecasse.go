package usecase

import (
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
)

type User = domain.User

type UserRepository interface {
	CreateUser(user User) (User, error)
	LoginUser(user User) (string, error)
	PromoteUser(userId string) (User, error)
}

type userUseCase struct {
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) domain.UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

// CreateUser inserts a new user into the database.
func (uuc *userUseCase) CreateUser(user User) (User, error) {
	return uuc.userRepository.CreateUser(user)

}

func (uuc *userUseCase) LoginUser(user User) (string, error) {
	return uuc.userRepository.LoginUser(user)

}

func (uuc *userUseCase) PromoteUser(userId string) (User, error) {
	return uuc.userRepository.PromoteUser(userId)
}
