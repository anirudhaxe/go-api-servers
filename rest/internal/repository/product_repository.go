package repository

import (
	"errors"
	"github.com/anirudhaxe/go-api-servers/rest/internal/model"
)

type ProductRepository struct {
	products map[string]*model.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]*model.Product),
	}
}

func (r *ProductRepository) Create(product *model.Product) error {

	if _, exists := r.products[product.ID]; exists {
		return errors.New("product already exists")
	}

	r.products[product.ID] = product

	return nil

}
func (r *ProductRepository) GetAll() []*model.Product {

	var productList []*model.Product

	for _, product := range r.products {
		productList = append(productList, product)
	}

	return productList

}
