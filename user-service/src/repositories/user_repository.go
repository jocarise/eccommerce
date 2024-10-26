package repositories

import (
	"github.com/jocarise/user-service/src/models"
)

type UserRepository interface {
	Create(user *models.User) error
}

type InMemoryUserRepository struct {
	users map[string]*models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]*models.User)}
}

func (r *InMemoryUserRepository) Create(user *models.User) error {
	r.users[user.Id] = user
	return nil
}
