package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	service services.UserUseCaseInterface
}

type UserControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	ListById(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewUserController(us services.UserUseCaseInterface) UserControllerInterface {
	return &UserController{
		service: us,
	}
}

func (us *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var u entities.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := us.service.Create(&u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
func (us *UserController) ListById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := us.service.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (us *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := us.service.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
