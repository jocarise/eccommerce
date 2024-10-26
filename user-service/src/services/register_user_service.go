package services

import (
	"errors"

	"github.com/jocarise/user-service/pkg/utils"
	"github.com/jocarise/user-service/src/models"
	"github.com/jocarise/user-service/src/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	if !util(user.Id) {
		return errors.New("invalid UUID")
	}

	if !utils.VerifyEmail(user.Email) {
		return errors.New("invalid email")
	}

	if !utils.VerifyPassword(user.Password) {
		return errors.New("invalid password")
	}

	return s.repo.Create(user)
}
