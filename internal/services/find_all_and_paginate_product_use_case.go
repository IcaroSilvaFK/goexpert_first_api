package services

import (
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
)

type FindAllAndPaginateProductUseCase struct {
	ProductDB *database.Product
}

type FindAllAndPaginateProductUseCaseInterface interface {
	Execute(page, limit int, sort string) (*[]entities.Product, error)
}

func NewFindAllAndPaginateProductUseCase(db *database.Product) FindAllAndPaginateProductUseCaseInterface {
	return &FindAllAndPaginateProductUseCase{
		ProductDB: db,
	}
}

func (pdb *FindAllAndPaginateProductUseCase) Execute(page, limit int, sort string) (*[]entities.Product, error) {

	products, err := pdb.ProductDB.FindAll(page, limit, sort)

	if err != nil {
		return nil, err
	}

	return products, nil
}
