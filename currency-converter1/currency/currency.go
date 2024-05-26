package currency

// Rates CurrencyRates - структура для хранения курсов валют.
//
// Base - Базовая валюта, по отношению к которой указаны курсы.
// Rates - Карта, содержащая курсы валют по отношению к базовой валюте. Ключом карты является код валюты, значением - курс.
type Rates struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}
