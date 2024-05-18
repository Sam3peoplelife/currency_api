package main

import (
    "log"
    "net/http"
)

//<summary>
//Creating server
//</summary>
func main() {
    http.HandleFunc("/", handleIndex)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}