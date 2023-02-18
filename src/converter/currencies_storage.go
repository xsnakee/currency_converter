package currency_converter

import "time"

type currencies_storage struct {
	currencies map[string]currency
	date       time.Time
	converter  converter
}

func (storage *currencies_storage) Convert(from string, to string, amount float64) float64 {
	currencyFrom := storage.currencies[from]
	currencyTo := storage.currencies[to]
	return storage.converter.Convert(currencyFrom, currencyTo, amount)
}
