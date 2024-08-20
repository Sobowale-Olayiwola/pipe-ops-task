package main

import (
	"testing"
	"time"
)

func TestSquare(t *testing.T) {
	t.Run("Square processes all numbers", func(t *testing.T) {
		done := make(chan struct{})
		genCh := make(chan int, 100)
		squareCh := make(chan int, 100)
		defer close(done)

		go func() {
			genCh <- 2
			genCh <- 3
			close(genCh)
		}()
		go square(done, genCh, squareCh)

		expected := []int{4, 9}
		i := 0
		for result := range squareCh {
			if result != expected[i] {
				t.Errorf("Expected %d, got %d", expected[i], result)
			}
			i++
		}
	})

	t.Run("Square stops early when done is called", func(t *testing.T) {
		done := make(chan struct{})
		genCh := make(chan int, 100)
		squareCh := make(chan int, 100)
		// Simulate early shutdown
		close(done)
		// Simulate a delay to ensure the value is processed before shutting down
		time.Sleep(10 * time.Millisecond)
		go func() {
			genCh <- 2
			genCh <- 3
			genCh <- 4
			close(genCh)
		}()
		go square(done, genCh, squareCh)

		// Check that no values are processed after done is closed
		for result := range squareCh {
			t.Errorf("Expected no result due to shutdown, but got %d", result)
		}
	})
}
