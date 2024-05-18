package main

import (
    "fmt"
    "html/template"
    //log"
    "net/http"
    "api/models"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
	"api/utils"
)

//<summary>
//Getting currency and executing it to html
//</summary>
func handleIndex(w http.ResponseWriter, r *http.Request) {
    rate, err := utils.FetchExchangeRate()
    if err != nil {
        http.Error(w, fmt.Sprintf("Could not fetch exchange rate: %v", err), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, fmt.Sprintf("Could not load template: %v", err), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, rate)
}

//<summary>
//Connection to DB and saving email
//</summary>
func handleSubscribe(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    email := r.FormValue("email")
    if email == "" {
        http.Error(w, "Email is required", http.StatusBadRequest)
        return
    }

    db, err := gorm.Open(sqlite.Open("exchange.db"), &gorm.Config{})
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }

    user := models.User{Email: email}
    if err := db.Create(&user).Error; err != nil {
        http.Error(w, "Could not save subscription", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Subscribed successfully")
}