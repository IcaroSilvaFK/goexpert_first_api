package services

import (
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
)

type FindProductByIdUseCase struct {
	ProductDB database.ProductInterface
}

type FindProductByIdUseCaseInterface interface {
	Execute(id string) (*entities.Product, error)
}

func NewFindProductByIdUseCase(db database.ProductInterface) FindProductByIdUseCaseInterface {
	return &FindProductByIdUseCase{
		ProductDB: db,
	}
}

func (db *FindProductByIdUseCase) Execute(id string) (*entities.Product, error) {

	p, err := db.ProductDB.FindById(id)

	if err != nil {
		return nil, err
	}

	return p, nil
}
