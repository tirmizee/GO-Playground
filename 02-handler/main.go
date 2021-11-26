package main

import (
	"app-handler/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/product", handlers.Product)
	http.HandleFunc("/customer", handlers.Customer)

	fmt.Println("server 8080 started")
	http.ListenAndServe(":8080", nil)
}
