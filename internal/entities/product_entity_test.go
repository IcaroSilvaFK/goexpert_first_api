package entities_test

import (
	"testing"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewProduct(t *testing.T) {

	p, err := entities.NewProduct("Test Product", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "Test Product", p.Name)
	assert.Equal(t, 10, p.Price)
	assert.NotNil(t, p.Price)
	assert.NotNil(t, p.CreatedAt)
}

func TestShouldNotCreateNewProductInvalidPriceLessThanZeroValue(t *testing.T) {

	p, err := entities.NewProduct("Test Product", -10)

	assert.Nil(t, p)
	assert.Error(t, err)
	assert.Equal(t, entities.ErrInvalidPrice, err)
}
func TestShouldNotCreateNewProductInvalidPriceZeroValue(t *testing.T) {

	p, err := entities.NewProduct("Test Product", 0)

	assert.Nil(t, p)
	assert.Error(t, err)
	assert.Equal(t, entities.ErrPriceIsRequired, err)
}

func TestShouldNotCreateNewProductInvalidName(t *testing.T) {

	p, err := entities.NewProduct("", -10)

	assert.Nil(t, p)
	assert.Error(t, err)
	assert.Equal(t, entities.ErrNameIsRequired, err)
}
