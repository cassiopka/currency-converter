# Конвертер Валют

Простой консольный конвертер валют, использующий API Open Exchange Rates для получения актуальных курсов.

## Установка

### Требования

* Go 1.22.3 или выше
* API ключ от Open Exchange Rates ([https://openexchangerates.org/](https://openexchangerates.org/))

### Инструкции

1. Клонируйте репозиторий проекта:
   ```bash
   git clone https://github.com/ваш_github/currency-converter.git
   ```
2. Перейдите в директорию проекта:
   ```bash
   cd currency-converter
   ```
3. Установите API ключ от Open Exchange Rates в переменную окружения `OPENEXCHANGERATES_API_KEY`:
   ```bash
   export OPENEXCHANGERATES_API_KEY="ваш_api_ключ"
   ```
4. Соберите приложение:
   ```bash
   go build
   ```

## Использование

1. Запустите приложение:
   ```bash
   ./currency-converter 
   ```
2. Введите код базовой валюты (например, USD).
3. Введите коды валют, для которых вы хотите узнать курсы, разделенные запятой (например, EUR,GBP,JPY).

**Пример:**

```
Введите код базовой валюты (например, USD): USD
Введите коды валют, разделенные запятой (например, EUR,GBP,JPY): EUR, GBP
Актуальные курсы валют:
------------------------
Базовая валюта: USD
------------------------
EUR: 0.9196
JPY: 156.9450
------------------------
```


## Разработка

### Архитектура

Проект реализован с использованием простой модульной архитектуры.  Он разделен на три пакета:

* `api`: Отвечает за взаимодействие с API Open Exchange Rates.
* `currency`: Содержит структуры данных для хранения информации о валютах.
* `ui`: Отвечает за взаимодействие с пользователем через консоль.

### Технологии

* Go 1.22.3
* Стандартная библиотека Go (`net/http`, `encoding/json`, `fmt`, `os`, `strings`)


