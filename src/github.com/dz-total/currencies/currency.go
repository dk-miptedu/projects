package currencies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// https://api.freecurrencyapi.com/v1/latest?apikey=%s&currencies=%s&base_currency=%s
const (
	//	apiURL = "https://freecurrencyapi.com/api/v2/latest?apikey=%s&base_currency=%s"
	apiURL = "api.freecurrencyapi.com/v1/latest?apikey=%s&currencies=%s&base_currency=%s"
)

type CurrencyClient struct {
	APIKey string
	Client *http.Client
}

func NewCurrencyClient(apiKey string) *CurrencyClient {
	return &CurrencyClient{
		APIKey: apiKey,
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *CurrencyClient) GetCurrencyRate(currencies string, baseCurrency string) (map[string]float64, error) {
	url := fmt.Sprintf(apiURL, c.APIKey, currencies, baseCurrency)
	fmt.Println(apiURL)
	fmt.Println(c.APIKey)
	fmt.Println(currencies)
	fmt.Println("baseCurrency Currency: ", baseCurrency)
	fmt.Println("url Currency: ", url)
	resp, err := c.Client.Get(url)
	fmt.Println("resp Currency: ", resp)
	if err != nil {
		return nil, fmt.Errorf("error making request to currency API: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d %s", resp.StatusCode, resp.Status)
	}
	var currencyResp CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&currencyResp); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return currencyResp.Data, nil
}
