package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello word")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contactez moi")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)

	fmt.Println("(http://localhost:8080)  -> server started on port", port)
	http.ListenAndServe(port, nil)
}
