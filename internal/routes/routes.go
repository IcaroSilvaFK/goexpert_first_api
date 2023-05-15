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
	pService := services.NewCreateProductUseCase(productDB)
	pController := controllers.NewProductController(pService)

	http.HandleFunc("/products", pController.Create)

}
