package main

import (
	"fmt"
	"go-crud/internal/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.Setup()

	fmt.Println("API em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
