package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Define a route for handling GET requests to the root path "/"
    r.HandleFunc("/", getHandler).Methods("GET")

    // Define a route for handling POST requests to "/post"
    r.HandleFunc("/post", postHandler).Methods("POST")

    port := 3000    
    fmt.Printf("Server is listening on :%d...\n", port)
    http.Handle("/", r)
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        panic(err)
    }
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    // Get the current working directory
    currentDir, err := os.Getwd()
    if err != nil {
        http.Error(w, "Failed to get current working directory", http.StatusInternalServerError)
        return
    }

    // Define the relative path to your HTML file
    relativePath := "main/index.html" // Modify this to your actual relative path

    // Create the absolute path by joining the current directory and the relative path
    absolutePath := filepath.Join(currentDir, relativePath)
	fmt.Println(absolutePath)
    // Serve the HTML file
    http.ServeFile(w, r, absolutePath)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    // This function will handle POST requests to "/post"
    fmt.Fprintf(w, "This is a POST request!")
}
