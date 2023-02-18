package currency_converter

type converter struct{}

func (converter *converter) Convert(from currency, to currency, amount float64) float64 {
	return float64(amount) / float64(from.rate) * float64(to.rate)
}
