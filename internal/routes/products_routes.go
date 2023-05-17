package routes

import (
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/controllers"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
	"github.com/go-chi/chi/v5"
)

func InitializeProductsRoutes(r *chi.Mux) {
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

	r.Post("/products", pController.Create)
	r.Get("/products", pController.List)
	r.Get("/products/{id}", pController.ListById)
	r.Put("/products/{id}", pController.Update)
	r.Delete("/products/{id}", pController.Update)
}
