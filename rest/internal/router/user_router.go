package router

import (
	"context"
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/handler"
	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
	"github.com/anirudhaxe/go-api-servers/rest/internal/service"
	"github.com/anirudhaxe/go-api-servers/rest/middleware"
	"github.com/anirudhaxe/go-api-servers/rest/utils"
)

func RegisterUserRoutes(ctx context.Context, db repository.DBTX) http.Handler {
	m := http.NewServeMux()

	r := repository.New(db)

	userService := service.NewUserService(r, ctx)
	userHandler := handler.NewUserHandler(userService)

	// m.HandleFunc("GET /product", utils.MakeHTTPHandlerFunc(productHandler.GetProducts))
	m.HandleFunc("POST /user", middleware.WithJWTAuth(utils.MakeHTTPHandlerFunc(userHandler.RegisterUser)))
	m.HandleFunc("POST /user/login", utils.MakeHTTPHandlerFunc(userHandler.LoginUser))

	return http.StripPrefix("/api/v1", m)

}
