package calc

import (
	"testing"
)

func TestAddition(t *testing.T) {
	expected := 3
	actual := Calculate(1, 2, add)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSubtraction(t *testing.T) {
	expected := -1
	actual := Calculate(2, 3, sub)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	expected := 6
	actual := Calculate(2, 3, mult)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDivision(t *testing.T) {
	expected := 2
	actual := Calculate(6, 3, div)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDivisionByZero(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected panic, got %v", err)
		}
	}()
	Calculate(2, 0, div)
}
