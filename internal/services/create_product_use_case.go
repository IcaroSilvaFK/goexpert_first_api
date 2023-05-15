package services

import (
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
)

type CreateProductUseCase struct {
	ProductDB database.ProductInterface
}

type CreateProductUseCaseInterface interface {
	Execute(*entities.Product) error
}

func NewCreateProductUseCase(pdb *database.Product) CreateProductUseCaseInterface {

	return &CreateProductUseCase{
		ProductDB: pdb,
	}

}

func (pdb *CreateProductUseCase) Execute(p *entities.Product) error {

	err := pdb.ProductDB.Create(p)

	if err != nil {
		return err
	}

	return nil
}
