package main

import (
	"log"
	"net/http"

	"github.com/juanplagos/bubble/router"
)

func main() {
	log.Println("server starting on ':8080'...")
	mux := router.RegisterRoutes()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
	    log.Fatal(err)
	}
}