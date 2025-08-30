package main

import (
	"fmt"
	"html/template" // Package for HTML templating
	"log"
	"net/http" // Package for HTTP server functionality
	"os"
	"os/signal"
	"syscall"
)

// Define a struct to hold data passed to the template
type PageData struct {
	Message string
}

// Global variable to store the parsed template for reuse
var tmpl *template.Template

// Initialize the template at startup
func init() {
	var err error
	tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error parsing template at startup: %v", err)
	}
}

// handler function to serve the HTML template
func handler(w http.ResponseWriter, r *http.Request) {
	// Create data to pass to the template
	data := PageData{
		Message: "Hello, World from Go!",
	}

	// Execute the template, writing the output to the http.ResponseWriter
	err := tmpl.Execute(w, data)
	if err != nil {
		// If there's an error executing the template, log it and send a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}

	fmt.Println("Successfully served index.html")
}

func main() {
	// Create a new HTTP server
	server := &http.Server{Addr: ":8080"}

	// Register the handler function for the root URL "/"
	http.HandleFunc("/", handler)

	// Start the server in a goroutine
	go func() {
		fmt.Println("Server starting on http://localhost:8080")
		log.Println("Server is running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("\nShutting down the server...")
	if err := server.Close(); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	fmt.Println("Server stopped.")
}
