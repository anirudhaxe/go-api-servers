package router

import (
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/handler"
	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
	"github.com/anirudhaxe/go-api-servers/rest/internal/service"
)

func RegisterProductRoutes() http.Handler {
	m := http.NewServeMux()

	repo := repository.NewProductRepository()
	productService := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(productService)

	m.HandleFunc("GET /", productHandler.GetProducts)
	m.HandleFunc("POST /", productHandler.CreateProduct)

	return http.StripPrefix("/product", m)

}
