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

	r.Route("/users", func(r chi.Router) {

		r.Post("/", uController.Create)
		r.Post("/token", uController.GetJWT)
		r.Get("/{id}", uController.ListById)
		r.Delete("/{id}", uController.Delete)
	})

}
