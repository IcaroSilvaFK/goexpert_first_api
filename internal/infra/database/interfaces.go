package database

import "github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(e string) (*entities.User, error)
	FindById(id string) (*entities.User, error)
	Delete(id string) error
}
