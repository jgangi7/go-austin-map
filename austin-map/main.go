package main

import (
    "austin-map/handlers"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    r := mux.NewRouter()

    // API routes
    r.HandleFunc("/api/locations", handlers.GetLocations).Methods("GET")

    // Serve static files
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}