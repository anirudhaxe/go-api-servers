package router

import (
	"context"
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/handler"
	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
	"github.com/anirudhaxe/go-api-servers/rest/internal/service"
	"github.com/anirudhaxe/go-api-servers/rest/utils"
)

func RegisterProductRoutes(ctx context.Context, db repository.DBTX) http.Handler {
	m := http.NewServeMux()

	r := repository.New(db)
	productService := service.NewProductService(r, ctx)
	productHandler := handler.NewProductHandler(productService)

	m.HandleFunc("GET /product", utils.MakeHTTPHandlerFunc(productHandler.GetProducts))
	m.HandleFunc("POST /product", utils.MakeHTTPHandlerFunc(productHandler.CreateProduct))

	return http.StripPrefix("/api/v1", m)

}
