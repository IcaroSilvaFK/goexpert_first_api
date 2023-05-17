package routes

import (
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/controllers"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func InitializeUserRoutes(r *chi.Mux, db *gorm.DB) {

	uDB := database.NewUserDB(db)

	userService := services.NewUserUseCase(uDB)
	uController := controllers.NewUserController(userService)

	r.Post("/users", uController.Create)
	r.Get("/users/{id}", uController.ListById)
	r.Delete("/users/{id}", uController.Delete)
}
