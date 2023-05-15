package database

import "github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(e string) (*entities.User, error)
	FindById(id string) (*entities.User, error)
	Delete(id string) error
}

type ProductInterface interface {
	Create(product *entities.Product) error
	FindAll(page, limit int, sort string) (*[]entities.Product, error)
	FindById(id string) (*entities.Product, error)
	Update(id, name string, price float64) error
	Delete(id string) error
}
