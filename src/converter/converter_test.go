package currency_converter

import (
	"reflect"
	"testing"
)

func TestCurrencyRateEquatable_MustReturnsTrue_ForSameValues(t *testing.T) {

	var val1 currencyRate = 123.00000001
	var val2 currencyRate = 123.00000001

	result, _ := val1.Equals(val2)

	if !result {
		t.Fatalf("%s equtable invalid, because values are same. %.10f != %.10f ", reflect.TypeOf(val1), val1, val2)
	}
}

func TestCurrencyRateEquatable_MustReturnsFalse_ForDifferentValues(t *testing.T) {

	var val1 currencyRate = 1111.00000001
	var val2 currencyRate = 1111.00000002

	result, _ := val1.Equals(val2)

	if result {
		t.Fatalf("%s equtable invalid, because values are different. %.10f == %.10f ", reflect.TypeOf(val1), val1, val2)
	}
}

func TestConvert_MustResturnsValidValue_ForSourceAndDestincationCurrencyAreSame(t *testing.T) {
	converter := converter{}
	expectedValue := 2.1
	src := currency{id: "USD", rate: 1}
	dest := src
	amount := expectedValue

	actualResult, err := converter.Convert(src, dest, amount)

	if expectedValue != actualResult || err != nil {
		t.Fatalf("Invalid convert. Expected value: %.10f. Actual value: %.10f. Value not same.", expectedValue, actualResult)
	}

}

func TestConvert_MustResturnsValidValue(t *testing.T) {
	converter := converter{}
	expectedValue := 3.0000000000
	src := currency{id: "RUB", rate: 70}
	dest := currency{id: "USD", rate: 1}
	amount := 210.0000

	actualResult, err := converter.Convert(src, dest, amount)

	if expectedValue != actualResult || err != nil {
		t.Fatalf("Invalid convert. Expected value: %.10f. Actual value: %.10f.", expectedValue, actualResult)
	}

}
