package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/anirudhaxe/go-api-servers/rest/internal/router"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	db, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Println(err.Error())
		log.Fatal("ERROR IN DB CONN")
	}
	defer db.Close(ctx)

	m := http.NewServeMux()

	userRoutes := router.RegisterUserRoutes(ctx, db)

	m.Handle("/api/v1/user", userRoutes)
	m.Handle("/api/v1/user/", userRoutes)

	log.Println("REST SERVER RUNNING ON: 3000")
	http.ListenAndServe(":3000", m)
}
