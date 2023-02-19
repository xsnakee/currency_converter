package currency_converter

type converter struct{}

func (converter *converter) Convert(from currency, to currency, amount float64) (float64, error) {
	// if from == nil || to == nil || math.from.rate || to.rate
	return float64(amount) / float64(from.rate) * float64(to.rate), nil
}
