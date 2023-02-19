package currency_converter

import (
	"errors"
	"math"
)

type converter struct{}

func (converter *converter) Convert(src currency, dest currency, amount float64) (float64, error) {

	srcRateIsZero, _ := src.rate.Equals(0)
	if srcRateIsZero {
		return 0, errors.New("src Rate is empty")
	}

	destRateIsZero, _ := dest.rate.Equals(0)
	if destRateIsZero {
		return 0, errors.New("dest Rate is empty")
	}

	return float64(dest.rate) * float64(amount) / float64(src.rate), nil
}

const float64EquantityThreshold = 1e-9

func (value currencyRate) Equals(otherValue currencyRate) (bool, error) {
	return equals(float64(value), float64(otherValue))
}

func equals(value float64, otherValue float64) (bool, error) {
	return math.Abs(value-otherValue) <= float64EquantityThreshold, nil
}
