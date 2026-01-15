package middleware

import (
	"fmt"
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/repository"
	"github.com/anirudhaxe/go-api-servers/rest/utils"
	"github.com/golang-jwt/jwt/v5"
)

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CALLING JWT AUTH MIDDLEWARE")

		tokenStr := r.Header.Get("x-jwt-token")

		token, err := repository.ValidateJwtToken(tokenStr)

		if err != nil {
			utils.PermissionDenied(w)
			return
		}

		if !token.Valid {
			utils.PermissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims["email"])

		// TODO: Can perform authentication match here by matching this claim email with second email in header (need to verify the existence of this email in the db as well). Alternatively, add this information to a context to then access it in the handlers

		handlerFunc(w, r)

	}
}
