package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
	"api/utils"

    "github.com/stretchr/testify/assert"
)

//<summary>
//Testing fetchExchangeRate function
//</summary>
func TestFetchExchangeRate(t *testing.T) {
    rate, err := utils.FetchExchangeRate()
    assert.NoError(t, err)
    assert.Greater(t, rate, 0.0, "The exchange rate should be greater than 0")
}

//<summary>
//Testing handleIndex function
//</summary>
func TestHandleIndex(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleIndex)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
    assert.Contains(t, rr.Body.String(), "Current USD to UAH Exchange Rate", "handler returned unexpected body")
}
