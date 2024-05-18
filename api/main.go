package main

import (
    "log"
    "net/http"
    //"api/models"
    "api/migrations"
    "api/email"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

//<summary>
//Creating server
//</summary>
func main() {
	db, err := gorm.Open(sqlite.Open("exchange.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

	migrations.Migrate(db)

	email.StartEmailScheduler(db)
	


    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/subscribe", handleSubscribe)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}