package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define handlers for different endpoints
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/api/data", apiDataHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page - Learn more about us!")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// For demonstration, just returning a static list of users
	users := []string{"Alice", "Bob", "Charlie"}
	fmt.Fprintf(w, "Users: %v", users)
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	// For demonstration, returning dummy JSON data
	jsonData := `{"name": "John", "age": 30, "city": "New York"}`
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, jsonData)
}
