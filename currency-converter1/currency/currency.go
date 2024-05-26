package currency

// Rates CurrencyRates - структура для хранения курсов валют
type Rates struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}
