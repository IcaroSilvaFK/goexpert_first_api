package routes

import (
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/controllers"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"
)

func InitializeProductsRoutes(r *chi.Mux, db *gorm.DB, jwt *jwtauth.JWTAuth) {

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
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt))
		r.Use(jwtauth.Authenticator)
		r.Post("/", pController.Create)
		r.Get("/", pController.List)
		r.Get("/{id}", pController.ListById)
		r.Put("/{id}", pController.Update)
		r.Delete("/{id}", pController.Delete)
	})

}
