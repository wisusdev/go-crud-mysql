package main

import (
	"log"
	"main/routes"
	"net/http"
)

func main() {
	routes.WebRoutes()

	port := ":8080"

	log.Printf("Server started on: http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		return
	}
}
