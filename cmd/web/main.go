package main

import (
	"fmt"
	"net/http"

	"main/internal/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Println("(http://localhost:8080)  -> server started on port", port)
	http.ListenAndServe(port, nil)
}
