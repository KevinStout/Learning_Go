package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		// Handle POST requests with JSON payload.
		var data map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		fmt.Println(decoder)
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		r.Body.Close()

		// convert the data map to a JSON string.
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the JSON string to the response.
		// If the file does not exist, Write creates. If the file exists, Write truncates (clear) it before writing.
		// 0644 is the file permission. This means the owner can read and write, and everyone else can only read.
		err = ioutil.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

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
