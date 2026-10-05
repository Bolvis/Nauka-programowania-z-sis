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
