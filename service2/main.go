package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from service 2")
	})
	http.HandleFunc("/users", usersHandler)

	http.HandleFunc("/users/{id}", userHandler)

	http.ListenAndServe(":8081", nil)
}
