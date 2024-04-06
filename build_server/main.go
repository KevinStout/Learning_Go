package main

import (
	"fmt"
	"net/http"

	"example.com/greetings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
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

	case "POST":
		// Handle POST request.
		// You can read data from the request body with r.Body.
		// For example, if the request body is JSON, you can decode it like this:
		// decoder := json.NewDecoder(r.Body)
		// err := decoder.Decode(&yourDataStructure)
		// Don't forget to handle the error and close the request body.
		// r.Body.Close()

	default:
		// Handle all other HTTP methods.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
