package main

import (
	"fmt"
	"net/http"

	"example.com/greetings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the "name" query param from the URL.
	name := r.URL.Query().Get("name")
	// If no name was given in the query param, use "Guest" as the default name.
	if name == "" {
		name = "Guest"

	}

	greeting, err := greetings.Hello(name)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	fmt.Fprintf(w, "%s", greeting)
	fmt.Fprintf(w, "Hello, Mom! %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
