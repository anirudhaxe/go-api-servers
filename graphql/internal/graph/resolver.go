package graph

import (
	"github.com/anirudhaxe/go-api-servers/graphql/internal/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	Repo *repository.Queries
}
