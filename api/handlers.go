package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "net/http"
)

const apiUrl = "https://api.exchangerate-api.com/v4/latest/USD"

type ExchangeRateResponse struct {
    Rates map[string]float64 `json:"rates"`
}

//<summary>
//Getting json file with USD currency
//<output>UAH currency</output>
//</summary>
func fetchExchangeRate() (float64, error) {
	//GET request
    resp, err := http.Get(apiUrl)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

	//checking status
    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("failed to fetch exchange rate: %v", resp.Status)
    }

	//Decoding json
    var rateResponse ExchangeRateResponse
    if err := json.NewDecoder(resp.Body).Decode(&rateResponse); err != nil {
        return 0, err
    }

	//Getting UAH rate
    rate, ok := rateResponse.Rates["UAH"]
    if !ok {
        return 0, fmt.Errorf("rate for UAH not found")
    }

    return rate, nil
}

//<summary>
//Getting currency and executing it to html
//</summary>
func handleIndex(w http.ResponseWriter, r *http.Request) {
    rate, err := fetchExchangeRate()
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
