package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/juanplagos/bubble/router"
)

func main() {
	mux := router.RegisterRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
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