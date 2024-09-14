package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Define a handler function for the root path
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to your Go server!")
    })

    // Start the server on port 8080
    fmt.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
