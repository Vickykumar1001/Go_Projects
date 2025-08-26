package repository

import (
	"database/sql"
	"errors"
	"user_management/internal/models"
)

type Repository struct {
	Conn *sql.Conn
}

func (i *Repository) GetAllUsers() ([]*models.User, error) {
	var slc []*models.User
	for _, v := range i.users {
		slc = append(slc, &v)

	}

	return slc, nil

}

func (i *InMemory) GetUserByUserName(userName string) (*models.User, error) {
	// We will
	if v, ok := i.users[userName]; ok {

		return &v, nil
	}
	return nil, errors.New("user Not Found")
}

func (i *InMemory) CreateUser(user models.User) string {
	// We wil create the user here
	i.users[user.Username] = user

	return user.ID
}
func (i *InMemory) UpdateUser(username string, user models.User) (models.User, error) {
	// We wil update the user here
	_, ok := i.users[username]
	if !ok {
		return models.User{}, errors.New("user Not Found")
	}

	i.users[username] = user

	return user, nil
}
func (i *InMemory) DeleteUser(username string) error {
	// We wil update the user here
	_, ok := i.users[username]
	if !ok {
		return errors.New("user Not Found")
	}

	delete(i.users, username)
	return nil
}

func NewInMemory() DbRepository {
	return &InMemory{
		users: make(map[string]models.User),
	}
}
