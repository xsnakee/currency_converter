package currency_converter

import "time"

type dataProvider interface {
	Load(time time.Time) currencies_storage
}
