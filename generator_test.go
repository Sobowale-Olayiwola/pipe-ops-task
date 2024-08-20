package main

import "testing"

// Mock implementation of RandIntn for testing
func mockRandIntn(n int) int {
	return n - 1 // Always return the maximum value for testing
}

func TestGenerator(t *testing.T) {
	t.Run("Generator completes all numbers", func(t *testing.T) {
		done := make(chan struct{})
		genCh := make(chan int, 100)
		defer close(done)

		go generator(mockRandIntn, done, 10, genCh)

		count := 0
		for range genCh {
			count++
		}

		if count != 10 {
			t.Errorf("Expected 10 generated numbers, got %d", count)
		}
	})

	t.Run("Generator stops early when done is called", func(t *testing.T) {
		done := make(chan struct{})
		genCh := make(chan int, 100)

		go generator(mockRandIntn, done, 1000, genCh)

		// Simulate an early shutdown
		go func() {
			// Close done after generating a few numbers
			// Simulate an early shutdown (before 1000 numbers are generated)
			close(done)
		}()

		count := 0
		for range genCh {
			count++
		}

		if count >= 1000 {
			t.Errorf("Expected generator to stop early, but generated %d numbers", count)
		} else {
			t.Logf("Generator stopped early as expected, generated %d numbers", count)
		}
	})
}
