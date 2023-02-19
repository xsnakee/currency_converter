package currency_converter

import "math"

const float64EquantityThreshold = 1e-9

func (value currencyRate) Equals(otherValue currencyRate) (bool, error) {
	return math.Abs(float64(value)-float64(otherValue)) <= float64EquantityThreshold, nil
}
