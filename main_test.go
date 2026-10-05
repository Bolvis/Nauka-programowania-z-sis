package main

import (
	"testing"
)

func TestCzyDawidMłodszy(t *testing.T) {
	actual := czyDawidMłodszy(24)
	expected := false
	if actual != expected {
		t.Errorf("The recieved value (%v) doesn't match expeded value (%v)", actual, expected)
	}
}

func TestOdmianaGramatycznaLat(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected string
	}{
		{name: "one", age: 1, expected: "rok"},
		{name: "two", age: 2, expected: "lata"},
		{name: "four", age: 4, expected: "lata"},
		{name: "five", age: 5, expected: "lat"},
		{name: "eleven", age: 11, expected: "lat"},
		{name: "fourteen", age: 14, expected: "lat"},
		{name: "twenty-one", age: 21, expected: "lat"},
		{name: "twenty-two", age: 22, expected: "lata"},
		{name: "twenty-five", age: 25, expected: "lat"},
		{name: "one-hundred", age: 100, expected: "lat"},
		{name: "one-hundred-one", age: 101, expected: "lat"},
		{name: "one-hundred-two", age: 102, expected: "lata"},
		{name: "one-hundred-four", age: 104, expected: "lata"},
		{name: "one-hundred-five", age: 105, expected: "lat"},
		{name: "one-hundred-eleven", age: 111, expected: "lat"},
		{name: "one-hundred-twenty-two", age: 122, expected: "lata"},
		{name: "maximum", age: 999, expected: "lat"},
		{name: "zero", age: 0, expected: "PODANA BŁĘDNA LICZBA"},
		{name: "negative", age: -1, expected: "PODANA BŁĘDNA LICZBA"},
		{name: "above-maximum", age: 1000, expected: "PODANA BŁĘDNA LICZBA"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := odmianaGramatycznaLat(test.age)
			if actual != test.expected {
				t.Errorf("odmianaGramatycznaLat(%d) = %q, want %q", test.age, actual, test.expected)
			}
		})
	}
}
