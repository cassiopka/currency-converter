package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// LatestRatesResponse Структура для хранения данных о курсах валют, полученных от Open Exchange Rates
type LatestRatesResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

// GetLatestRates получает актуальные курсы валют с Open Exchange Rates.
// Возвращает структуру LatestRatesResponse и ошибку (при наличии).
func GetLatestRates(apiKey string) (*LatestRatesResponse, error) {
	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении данных: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неуспешный статус код: %s", resp.Status)
	}

	var ratesResponse LatestRatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&ratesResponse); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании JSON: %w", err)
	}

	return &ratesResponse, nil
}

// GetApiKey получает API ключ из переменных окружения.
// Возвращает API ключ и флаг, указывающий на наличие ключа.
func GetApiKey() (string, bool) {
	apiKey := os.Getenv("OPENEXCHANGERATES_API_KEY")
	return apiKey, apiKey != ""
}
