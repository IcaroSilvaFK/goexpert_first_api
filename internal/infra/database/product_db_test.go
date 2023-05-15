package database_test

import (
	"testing"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func generateDatabase() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.Product{})
	return db
}

func TestShouldCreateNewProduct(t *testing.T) {

	db := generateDatabase()

	productDB := database.NewProductDB(db)

	p, _ := entities.NewProduct("Product 1", 10)

	err := productDB.Create(p)

	assert.Nil(t, err)

	var product entities.Product

	err = db.Where("id = ?", p.ID).First(&product).Error

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, p.Name, product.Name)
	assert.Equal(t, p.ID, product.ID)
	assert.Equal(t, p.Price, product.Price)
}

func TestShouldFindProductById(t *testing.T) {

	db := generateDatabase()

	productDB := database.NewProductDB(db)

	p, _ := entities.NewProduct("Product 1", 10)

	err := productDB.Create(p)

	assert.Nil(t, err)

	product, err := productDB.FindById(p.ID.String())

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, p.Name, product.Name)
	assert.Equal(t, p.ID, product.ID)
	assert.Equal(t, p.Price, product.Price)
}

func TestShouldUpdateProduct(t *testing.T) {

	db := generateDatabase()

	productDB := database.NewProductDB(db)

	p, _ := entities.NewProduct("Product 1", 10)

	err := productDB.Create(p)

	assert.Nil(t, err)

	err = productDB.Update(p.ID.String(), "Teste", 20)

	assert.Nil(t, err)

	product, err := productDB.FindById(p.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.Name, "Teste")
	assert.Equal(t, product.ID, product.ID)
	assert.Equal(t, product.Price, float64(20))
}

func TestShouldListAllProducts(t *testing.T) {

	db := generateDatabase()

	productDB := database.NewProductDB(db)

	p, _ := entities.NewProduct("Product 1", 10)
	p1, _ := entities.NewProduct("Product 2", 10)
	p2, _ := entities.NewProduct("Product 3", 10)
	p3, _ := entities.NewProduct("Product 4", 10)

	db.CreateInBatches([]*entities.Product{p, p1, p2, p3}, 3)

	err := productDB.Update(p.ID.String(), "Teste", 20)

	assert.Nil(t, err)

	product, err := productDB.ListAll(0, 0, "asc")

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Len(t, *product, 4)

}

func TestShouldListAllProductsByPage(t *testing.T) {

	db := generateDatabase()

	productDB := database.NewProductDB(db)

	p, _ := entities.NewProduct("Product 1", 10)
	p1, _ := entities.NewProduct("Product 2", 10)
	p2, _ := entities.NewProduct("Product 3", 10)
	p3, _ := entities.NewProduct("Product 4", 10)

	db.CreateInBatches([]*entities.Product{p, p1, p2, p3}, 3)

	product, err := productDB.ListAll(1, 2, "asc")

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Len(t, *product, 2)

}

func TestShouldDeleteProduct(t *testing.T) {

	db := generateDatabase()

	productDB := database.NewProductDB(db)

	p, _ := entities.NewProduct("Product 1", 10)

	err := productDB.Create(p)

	assert.Nil(t, err)

	err = productDB.Delete(p.ID.String())

	assert.Nil(t, err)

	product, err := productDB.FindById(p.ID.String())
	assert.Error(t, err)
	assert.Nil(t, product)

}
