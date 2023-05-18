package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/dtos"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type UserController struct {
	service  services.UserUseCaseInterface
	jwt      *jwtauth.JWTAuth
	jwtExpIn int
}

type UserControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	ListById(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetJWT(w http.ResponseWriter, r *http.Request)
}

func NewUserController(us services.UserUseCaseInterface, jwt *jwtauth.JWTAuth, jwtExpIn int) UserControllerInterface {
	return &UserController{
		service:  us,
		jwt:      jwt,
		jwtExpIn: jwtExpIn,
	}
}

func (us *UserController) GetJWT(w http.ResponseWriter, r *http.Request) {

	var user dtos.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := us.service.FindByEmail(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	fmt.Println(u)

	if err = u.ValidatePassword(user.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized email or password is invalid"))
		return
	}
	fmt.Println(us.jwtExpIn)
	_, tokenString, _ := us.jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(us.jwtExpIn)).Unix(),
	})

	accessTokenPayload := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessTokenPayload)
}

func (us *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user dtos.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entities.NewUserEntity(user.Email, user.Name, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	err = us.service.Create(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// secret,err := configs.LoadConfig(".")

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: " + err.Error()))
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
