package main

import (
	"api/platzi/handler"
	"api/platzi/server"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error", err)
	}

	port := os.Getenv("PORT")
	JwtSecert := os.Getenv("JWT_SECRET")
	dbUrl := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JwtSecert,
		Port:        port,
		DataBaseUrl: dbUrl,
	})

	s.Start(bindRoutes)
}

func bindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handler.HomeHandler(s)).Methods(http.MethodGet)
}
