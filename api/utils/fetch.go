package utils

import (
    "encoding/json"
    "fmt"
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
func FetchExchangeRate() (float64, error) {
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