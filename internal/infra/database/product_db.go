package database

import (
	"strings"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProductDB(conn *gorm.DB) *Product {
	return &Product{
		DB: conn,
	}
}

func (p *Product) Create(product *entities.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) (*[]entities.Product, error) {

	if strings.TrimSpace(sort) != "" && sort != "asc" && sort != "desc" { //O(1)
		sort = "asc"
	}

	var products []entities.Product
	var err error

	current := (page - 1) * limit //O(1)
	if page != 0 && limit != 0 {  //O(1)
		err = p.DB.Limit(limit).Offset(current).Find(&products).Order("created_at" + sort).Error

	} else {
		err = p.DB.Find(&products).Order("created_at" + sort).Error //O(n)
	}

	return &products, err
}

func (p *Product) FindById(id string) (*entities.Product, error) {

	var product entities.Product

	err := p.DB.Where("id = ?", id).First(&product).Error //O(n)

	if err != nil { //O(1)
		return nil, err
	}

	return &product, nil
}

func (p *Product) Update(id, name string, price float64) error {

	var product entities.Product

	tx := p.DB.Where("id = ?", id).First(&product)

	if tx.Error != nil {
		return tx.Error
	}

	product.Name = name
	product.Price = price

	return p.DB.Save(&product).Error

}

func (p *Product) Delete(id string) error {

	tx := p.DB.Where("id = ?", id)

	if tx.Error != nil {
		return tx.Error
	}

	err := tx.Delete(&entities.Product{}).Error

	return err
}
