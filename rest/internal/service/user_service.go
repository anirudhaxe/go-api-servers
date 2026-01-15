package service

import (
	"context"

	"github.com/anirudhaxe/go-api-servers/rest/internal/model"
	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// The service layer binds the repo to the handler layer
type UserService struct {
	repo *repository.Queries
	ctx  context.Context
}

func NewUserService(repo *repository.Queries, ctx context.Context) *UserService {
	return &UserService{repo: repo, ctx: ctx}
}

// registers a new user and returns token
func (s *UserService) RegisterUser(p *model.RegisterUserRequest) (*repository.User, error) {

	uuidStr := uuid.NewString()
	var pgUUID pgtype.UUID

	err := pgUUID.Scan(uuidStr)

	if err != nil {
		return nil, err
	}

	encpw, err := repository.GeneratePasswordHash(p.Password)

	if err != nil {
		return nil, err
	}

	pcp := repository.CreateUserParams{ID: pgUUID, Username: p.Username, Email: p.Email, EncryptedPassword: encpw, Role: p.Role, IsActive: true}

	usr, err := s.repo.CreateUser(s.ctx, pcp)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (s *UserService) GetUser(email string) (*repository.User, error) {

	usr, err := s.repo.GetUser(s.ctx, email)

	if err != nil {
		return nil, err
	}

	return &usr, nil

}

// // Will return the jwt token
// func (s *UserService) Login() {}
//
// func (s *UserService) ListProducts() ([]model.Product, error) {
// 	repoProducts, err := s.repo.GetAllProducts(s.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return mapRepoProductsToModel(repoProducts), nil
// }
//
// func mapRepoProductsToModel(repoProducts []repository.Product) []model.Product {
// 	products := make([]model.Product, len(repoProducts))
// 	for i, p := range repoProducts {
// 		products[i] = model.Product{
// 			ID:          p.ID,
// 			Name:        p.Name,
// 			Description: p.Description.String,
// 			Price:       p.Price.Float64,
// 		}
// 	}
// 	return products
// }
