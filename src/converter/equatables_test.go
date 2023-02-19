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
