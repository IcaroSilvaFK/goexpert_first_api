package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
)

type ProductController struct {
	createProductService   services.CreateProductUseCaseInterface
	findByIdProductService services.FindProductByIdUseCaseInterface
	findAllProductService  services.FindAllAndPaginateProductUseCaseInterface
	deleteProductService   services.DeleteProductUseCaseInterface
}

type ProductControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	// Update(w http.ResponseWriter, r *http.Request)
	// List(w http.ResponseWriter, r *http.Request)
	// Delete(w http.ResponseWriter, r *http.Request)
}

func NewProductController(
	createProductService services.CreateProductUseCaseInterface,
	findByIdProductService services.FindProductByIdUseCaseInterface,
	findAllProductService services.FindAllAndPaginateProductUseCaseInterface,
	deleteProductService services.DeleteProductUseCaseInterface,
) ProductControllerInterface {

	return &ProductController{
		createProductService:   createProductService,
		findByIdProductService: findByIdProductService,
		findAllProductService:  findAllProductService,
		deleteProductService:   deleteProductService,
	}

}

func (ct *ProductController) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}

	var p entities.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := ct.createProductService.Execute(&p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// func (ct *ProductController) Update(w http.ResponseWriter, r *http.Request)
// func (ct *ProductController) List(w http.ResponseWriter, r *http.Request)
// func (ct *ProductController) Delete(w http.ResponseWriter, r *http.Request)
