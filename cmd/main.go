package main

import (
	"log"
	"net/http"
	"os"

	"github.com/juanplagos/bubble/repository"
	"github.com/juanplagos/bubble/router"
	"github.com/rs/cors"
)

func main() {
	pool := repository.InitPostgresPool()
	defer pool.Close()
	mux := router.RegisterRoutes(pool)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ALLOWED_ORIGIN")},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
        AllowCredentials: true,
	})

	handler := c.Handler(mux)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
	    log.Fatal(err)
	}
}