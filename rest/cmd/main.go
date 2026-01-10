package main

import (
	"log"
	"net/http"

	"github.com/anirudhaxe/go-api-servers/rest/internal/router"
)

func main() {
	log.Println("cmd main.go file")

	m := http.NewServeMux()

	productRoutes := router.RegisterProductRoutes()

	m.Handle("/product/", productRoutes)

	log.Println("REST SERVER RUNNING ON: 3000")
	http.ListenAndServe(":3000", m)
}
