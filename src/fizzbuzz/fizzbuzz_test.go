package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	value := 3
	expected := "fizz"
	result := Fizzbuzz(value)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}
	value = 30
	expected = "fizz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}

	value = 15
	expected = "fizz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}
	value = 21
	expected = "fizz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}
	value = 99
	expected = "fizz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}

	value = 5
	expected = "buzz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}

	value = 25
	expected = "buzz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}

	value = 500
	expected = "buzz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}

	value = 100
	expected = "buzz"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}

	value = 2
	expected = "2"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}

	value = 99
	expected = "99"
	result = Fizzbuzz(value)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v is expected but got %v", expected, result)
	}
}
