package main

import (
	"log"
	"net/http"
	"Todo/routes"
)

func main() {
	r := routes.RegisterRoutes()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
