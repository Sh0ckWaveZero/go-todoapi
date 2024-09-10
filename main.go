package main

import (
	"fmt"
	"net/http"
	"todoapi/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
