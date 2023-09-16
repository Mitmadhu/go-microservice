package main

import (
	"fmt"
	"net/http"
"os"
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
    filePath := "static/index.html"

    // Use os.Stat to get file information
    _, err := os.Stat(filePath)

    if err == nil {
        // The file exists
        fmt.Printf("File %s exists.\n", filePath)
    } else if os.IsNotExist(err) {
        // The file does not exist
        fmt.Printf("File %s does not exist.\n", filePath)
    } else {
        // An error occurred while checking the file
        fmt.Printf("Error checking file %s: %v\n", filePath, err)
    }
    
    http.ServeFile(w, r, "static/index.html")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    // This function will handle POST requests to "/post"
    fmt.Fprintf(w, "This is a POST request!")
}
