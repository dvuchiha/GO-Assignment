package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const COINDESK_API_URL = "https://api.coindesk.com/v1/bpi/currentprice.json"

// Fetches response & parses json
func fetchData(apiURL string) (map[string]string, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var json_response map[string]interface{}
	unmarshal_err := json.Unmarshal(body, &json_response)

	if unmarshal_err != nil {
		return nil, unmarshal_err
	}
	pricesData := extractPrices(json_response)

	return pricesData, nil

}

// Extract prices of bitcoins from the coindesk-API response
func extractPrices(response map[string]interface{}) map[string]string {

	bitcoin_prices := map[string]string{}

	bpiMap := response["bpi"].(map[string]interface{})

	for country_code, items := range bpiMap {
		item_map := items.(map[string]interface{})
		bitcoin_prices[country_code] = item_map["rate"].(string)
	}

	return bitcoin_prices
}
