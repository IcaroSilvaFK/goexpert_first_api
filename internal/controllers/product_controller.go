package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/services"
	"github.com/go-chi/chi/v5"
)

type Input struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductController struct {
	createProductService   services.CreateProductUseCaseInterface
	findByIdProductService services.FindProductByIdUseCaseInterface
	findAllProductService  services.FindAllAndPaginateProductUseCaseInterface
	deleteProductService   services.DeleteProductUseCaseInterface
	updateProductService   services.UpdateProductUseCaseInterface
}

type ProductControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	ListById(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewProductController(
	createProductService services.CreateProductUseCaseInterface,
	findByIdProductService services.FindProductByIdUseCaseInterface,
	findAllProductService services.FindAllAndPaginateProductUseCaseInterface,
	deleteProductService services.DeleteProductUseCaseInterface,
	updateProductService services.UpdateProductUseCaseInterface,
) ProductControllerInterface {

	return &ProductController{
		createProductService:   createProductService,
		findByIdProductService: findByIdProductService,
		findAllProductService:  findAllProductService,
		deleteProductService:   deleteProductService,
		updateProductService:   updateProductService,
	}
}

func (ct *ProductController) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	var i Input

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input" + err.Error()))
		return
	}

	p, err := entities.NewProduct(i.Name, i.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input" + err.Error()))
		return
	}
	err = ct.createProductService.Execute(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on create product " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (ct *ProductController) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}

	var p entities.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := ct.updateProductService.Execute(id, p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on update product " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ct *ProductController) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	var pageInt, limitInt int
	var err error

	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Expected a valid number"))
			return
		}
	}
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Expected a valid number"))
			return
		}
	}

	products, err := ct.findAllProductService.Execute(pageInt, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error" + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

func (ct *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}

	err := ct.deleteProductService.Execute(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on delete product" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ct *ProductController) ListById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is missing a follow property"))
		return
	}

	p, err := ct.findByIdProductService.Execute(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error executing the service" + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}
