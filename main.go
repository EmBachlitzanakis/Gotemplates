package main

import (
	"fmt"
	"html/template" // Package for HTML templating
	"log"
	"net/http" // Package for HTTP server functionality
)

// Define a struct to hold data passed to the template
type PageData struct {
	Message string
}

// handler function to serve the HTML template
func handler(w http.ResponseWriter, r *http.Request) {
	// Parse the template file.
	// In a real application, you might parse templates once at startup
	// for better performance, but for "Hello World", parsing on each
	// request is fine for simplicity.
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		// If there's an error parsing the template, log it and send a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	// Create data to pass to the template
	data := PageData{
		Message: "Hello, World from Go!",
	}

	// Execute the template, writing the output to the http.ResponseWriter
	err = tmpl.Execute(w, data)
	if err != nil {
		// If there's an error executing the template, log it and send a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}

	fmt.Println("Successfully served index.html")
}

func main() {
	// Register the handler function for the root URL "/"
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	// log.Fatal will print any error and then exit the program
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
