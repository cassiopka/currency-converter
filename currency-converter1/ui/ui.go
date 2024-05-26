package ui

import (
	"fmt"
	"os"
	"strings"

	"currency-converter/currency"
)

// DisplayRates отображает курсы валют в консоли.
//
// Принимает список курсов валют и базовую валюту.
func DisplayRates(rates *currency.Rates) {
	fmt.Println("Актуальные курсы валют:")
	fmt.Println("------------------------")
	fmt.Printf("Базовая валюта: %s\n", rates.Base)
	fmt.Println("------------------------")
	for code, rate := range rates.Rates {
		fmt.Printf("%s: %.4f\n", code, rate)
	}
	fmt.Println("------------------------")
}

// GetBaseCurrency запрашивает у пользователя базовую валюту.
//
// Возвращает код базовой валюты (трехбуквенный код).
func GetBaseCurrency() string {
	var baseCurrency string
	for {
		fmt.Print("Введите код базовой валюты (например, USD): ")
		_, err := fmt.Scanln(&baseCurrency)
		if err != nil {
			return ""
		}
		baseCurrency = strings.ToUpper(baseCurrency)
		if len(baseCurrency) == 3 {
			return baseCurrency
		}
		fmt.Println("Неверный код валюты. Пожалуйста, введите трехбуквенный код.")
	}
}

// GetTargetCurrencies запрашивает у пользователя список целевых валют.
//
// Возвращает слайс строк с кодами валют.
func GetTargetCurrencies() []string {
	var currenciesStr string
	fmt.Print("Введите коды валют, разделенные запятой (например, EUR,GBP,JPY): ")
	_, err := fmt.Scanln(&currenciesStr)
	if err != nil {
		return nil
	}
	currencies := strings.Split(currenciesStr, ",")
	for i := range currencies {
		currencies[i] = strings.ToUpper(strings.TrimSpace(currencies[i])) // <-- Добавляем TrimSpace
	}
	return currencies
}

// ShowError выводит сообщение об ошибке в консоль.
func ShowError(err error) {
	_, err = fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
	if err != nil {
		return
	}
}
