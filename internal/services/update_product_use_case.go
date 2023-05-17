package services

import (
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
)

type UpdateProductUseCase struct {
	productDB database.ProductInterface
}

type UpdateProductUseCaseInterface interface {
	Execute(id string, p entities.Product) error
}

func NewUpdateProductUseCase(pDB database.ProductInterface) UpdateProductUseCaseInterface {
	return &UpdateProductUseCase{
		productDB: pDB,
	}
}

func (pdb *UpdateProductUseCase) Execute(id string, p entities.Product) error {

	err := pdb.productDB.Update(id, p.Name, p.Price)

	return err

}
