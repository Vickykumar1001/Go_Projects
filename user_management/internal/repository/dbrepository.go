package repository

import "user_management/internal/models"

type DbRepository interface {
	CreateUser(user models.User) string
	GetUserByUserName(userName string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(username string, user models.User) (models.User, error)
	DeleteUser(username string) error
}
