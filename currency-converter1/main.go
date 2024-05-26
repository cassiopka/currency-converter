package main

import (
	"currency-converter/api"
	"currency-converter/currency"
	"currency-converter/ui"
	"fmt"
	"os"
)

// main - Точка входа в приложение.
//
// Эта функция выполняет следующие действия:
//  1. Проверяет наличие API-ключа в переменной окружения OPENEXCHANGERATES_API_KEY.
//  2. Если ключ не найден, выводит сообщение об ошибке и завершает работу приложения.
//  3. Получает актуальные курсы валют с API Open Exchange Rates.
//  4. Обрабатывает ошибки при получении курсов.
//  5. Получает от пользователя базовую валюту и список валют, для которых необходимо получить курсы.
//  6. Создает структуру Rates с полученными курсами.
//  7. Отображает курсы валют пользователю.
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
