package routes

import (
	"net/http"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/controllers"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
)

func InitializeRoutes() {
	db := database.InitializeDatabase()
	db.AutoMigrate(&entities.User{}, &entities.Product{})

	productDB := database.NewProductDB(db)
	pServiceCreateProductService := services.NewCreateProductUseCase(productDB)
	pServiceFindByIdProductService := services.
		NewFindProductByIdUseCase(productDB)
	pServiceFindAllProductService := services.
		NewFindAllAndPaginateProductUseCase(productDB)
	pServiceUpdateProductService := services.NewUpdateProductUseCase(productDB)
	pServiceDeleteProductService := services.NewDeleteProductUseCase(productDB)

	pController := controllers.NewProductController(
		pServiceCreateProductService,
		pServiceFindByIdProductService,
		pServiceFindAllProductService,
		pServiceDeleteProductService,
		pServiceUpdateProductService,
	)

	http.HandleFunc("/products", pController.Create)
	http.HandleFunc("/products", pController.Create)

}
