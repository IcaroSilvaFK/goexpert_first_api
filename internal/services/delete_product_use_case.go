package services

import "github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"

type DeleteProductUseCase struct {
	ProductDB *database.Product
}

type DeleteProductUseCaseInterface interface {
	Execute(id string) error
}

func NewDeleteProductUseCase(db *database.Product) DeleteProductUseCaseInterface {
	return &DeleteProductUseCase{
		ProductDB: db,
	}
}

func (pdb *DeleteProductUseCase) Execute(id string) error {

	err := pdb.ProductDB.Delete(id)

	return err
}
