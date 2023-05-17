package main

import (
	"log"
	"net/http"

	"github.com/IcaroSilvaFK/goexpert_first_api/configs"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/routes"
	"github.com/go-chi/chi/v5"
)

func main() {

	cfg, err := configs.LoadConfig(".")
	r := chi.NewRouter()

	if err != nil {
		log.Fatal(err)
	}

	routes.InitializeProductsRoutes(r)
	routes.InitializeUserRoutes(r)

	log.Println("🚀Server running at port", cfg.WebServerPort)

	http.ListenAndServe(cfg.WebServerPort, r)
}

// type ProductHandler struct {
// 	ProductDB database.ProductInterface
// }

// func NewProductHandler(db database.ProductInterface) *ProductHandler {
// 	return &ProductHandler{
// 		ProductDB: db,
// 	}
// }

// func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	var product dtos.CreateProductInput

// 	err := json.NewDecoder(r.Body).Decode(&product)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	p, err := entities.NewProduct(product.Name, product.Price)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = h.ProductDB.Create(p)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)

// }
