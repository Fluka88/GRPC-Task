package main

import (
	"fmt"

	"A_Service/controller"
	"A_Service/middleware"
	"A_Service/model"
	"log"
	"net/http"
)

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

func main() {
	model.SetCache()
	mux := controller.Register()
	model.Connect()
	defer model.Close()
	fmt.Printf("Serving...")

	handler := middleware.Logging(mux)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
