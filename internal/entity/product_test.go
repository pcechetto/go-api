package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("test", 10.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, 10.00, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10.00)
	assert.Equal(t, ErrNameIsRequired, err)
	assert.Nil(t, product)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("test", 0.0)
	assert.Equal(t, ErrPriceIsRequired, err)
	assert.Nil(t, product)
}

func TestProductWhenPriceMustBeGreaterThanZero(t *testing.T) {
	product, err := NewProduct("test", -10.00)
	assert.Equal(t, ErrPriceMustBeGreaterThanZero, err)
	assert.Nil(t, product)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("test", 10.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
