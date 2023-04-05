package main

import (
	"fmt"
	"net/http"
)

func main() {
	initDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from service 1")
	})

	http.HandleFunc("/users", usersHandler)

	http.ListenAndServe(":8080", nil)
}
