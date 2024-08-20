package main

import (
	"os"
	"testing"
)

func TestRunPipeline(t *testing.T) {
	// Control the random number generation to return sequential numbers
	mockRandIntn := func(n int) int {
		return n - 1 // Always return the maximum value for testing
	}
	// Test with a small number of generated integers for quick testing
	numCount := 10

	// Run the pipeline
	finalSum := runPipeline(mockRandIntn, numCount)

	// Calculate the expected sum manually, since random will always return max which is 100
	// square of 100 = 10_000, and generate 10 numbers, 10_000 * 10 = 100_000
	expectedSum := (maxRandNum * maxRandNum) * 10

	// Assert the final sum is as expected
	if finalSum != expectedSum {
		t.Errorf("Expected sum %d, but got %d", expectedSum, finalSum)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
