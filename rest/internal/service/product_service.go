package service

import (
	"context"

	"github.com/anirudhaxe/go-api-servers/rest/internal/model"
	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

// The service layer binds the repo to the handler layer
type ProductService struct {
	repo *repository.Queries
	ctx  context.Context
}

func NewProductService(repo *repository.Queries, ctx context.Context) *ProductService {
	return &ProductService{repo: repo, ctx: ctx}
}

func (s *ProductService) CreateProduct(p *model.Product) error {

	pcp := repository.CreateProductParams{ID: p.ID, Name: p.Name, Description: pgtype.Text{String: p.Description, Valid: true}, Price: pgtype.Float8{Float64: p.Price, Valid: true}}

	err := s.repo.CreateProduct(s.ctx, pcp)

	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) ListProducts() ([]model.Product, error) {
	repoProducts, err := s.repo.GetAllProducts(s.ctx)
	if err != nil {
		return nil, err
	}

	return mapRepoProductsToModel(repoProducts), nil
}

func mapRepoProductsToModel(repoProducts []repository.Product) []model.Product {
	products := make([]model.Product, len(repoProducts))
	for i, p := range repoProducts {
		products[i] = model.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description.String,
			Price:       p.Price.Float64,
		}
	}
	return products
}
