package service

import (
	"github.com/anirudhaxe/go-api-servers/rest/internal/model"
	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
)

// The service layer binds the repo to the handler layer
type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {

	return &ProductService{repo: repo}

}

func (s *ProductService) AddProduct(product *model.Product) error {
	return s.repo.Create(product)
}
func (s *ProductService) ListProducts() []*model.Product {
	return s.repo.GetAll()
}
