package main

import (
	"testing"
)

func TestRandomNumberInRange(t *testing.T) {
	// Mock rand.Intn to return a predictable result
	mockRandIntn := func(n int) int {
		return n - 1 // Always return the maximum possible value for testing
	}

	// Test case 1: min = 1, max = 100, expect max value
	min, max := 1, 100
	expected := max
	result := randomNumberInRange(mockRandIntn, min, max)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: min = 10, max = 10, expect exactly 10
	min, max = 10, 10
	expected = 10
	result = randomNumberInRange(mockRandIntn, min, max)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 3: min > max, should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when min > max, but function did not panic")
		}
	}()
	randomNumberInRange(mockRandIntn, 100, 10)
}
