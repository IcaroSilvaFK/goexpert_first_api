package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/IcaroSilvaFK/goexpert_first_api/configs"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/time/rate"
)

type RateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.Mutex
	r   rate.Limit
	b   int
}

func NewIPRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.Mutex{},
		r:   r,
		b:   b,
	}
}

func (r *RateLimiter) AddIp(ip string) *rate.Limiter {

	r.mu.Lock()
	defer r.mu.Unlock()

	limiter := rate.NewLimiter(r.r, r.b)

	r.ips[ip] = limiter

	return limiter
}

func (r *RateLimiter) GetLimiter(ip string) *rate.Limiter {

	r.mu.Lock()
	defer r.mu.Unlock()

	limiter, exists := r.ips[ip]

	if !exists {
		r.mu.Unlock()
		return r.AddIp(ip)
	}

	r.mu.Unlock()

	return limiter
}

var limiter = NewIPRateLimiter(1, 5)

func limitMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			limiter := limiter.GetLimiter(r.RemoteAddr)

			if !limiter.Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		},
	)

}

func main() {

	cfg, err := configs.LoadConfig(".")
	r := chi.NewRouter()
	db := database.InitializeDatabase()
	db.AutoMigrate(&entities.User{}, &entities.Product{})

	if err != nil {
		log.Fatal(err)
	}

	r.Use(middleware.Logger)
	routes.InitializeProductsRoutes(r, db, cfg.TokenAuth)
	routes.InitializeUserRoutes(r, db, cfg.TokenAuth, cfg.JWTExpiresIn)

	log.Println("🚀Server running at port", cfg.WebServerPort)

	http.ListenAndServe(cfg.WebServerPort, limitMiddleware(r))
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
