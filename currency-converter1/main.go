package main

import (
	"fmt"
	"os"

	"currency-converter/api"
	"currency-converter/currency"
	"currency-converter/ui"
)

func main() {
	apiKey, exists := api.GetApiKey()
	if !exists {
		ui.ShowError(fmt.Errorf("api ключ не найден. установите переменную окружения OPENEXCHANGERATES_API_KEY"))
		os.Exit(1)
	}

	latestRates, err := api.GetLatestRates(apiKey)
	if err != nil {
		ui.ShowError(err)
		os.Exit(1)
	}

	baseCurrency := ui.GetBaseCurrency()
	targetCurrencies := ui.GetTargetCurrencies()

	rates := currency.Rates{
		Base:  baseCurrency,
		Rates: make(map[string]float64),
	}
	for _, targetCurrency := range targetCurrencies {
		if rate, ok := latestRates.Rates[targetCurrency]; ok {
			rates.Rates[targetCurrency] = rate
		} else {
			ui.ShowError(fmt.Errorf("валюта %s не найдена", targetCurrency))
		}
	}

	ui.DisplayRates(&rates)
}
