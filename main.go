package main

import (
	"log"
	"net/http"
	"Todo/routes"
	"Todo/config"
)

func main() {

config.LoadEnv()

	port := config.GetEnv("PORT", "8080")
	r := routes.RegisterRoutes()

	log.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
