package main

import (
	"testing"
	"time"
)

func TestSum_NormalOperation(t *testing.T) {
	// Create a done channel (it will remain open in this test)
	done := make(chan struct{})

	// Create a squareCh channel and populate it with squared numbers
	squareCh := make(chan int, 3)
	squareCh <- 4  // 2*2
	squareCh <- 9  // 3*3
	squareCh <- 16 // 4*4
	close(squareCh)

	resultCh := sum(done, squareCh)
	result := <-resultCh

	// Expected sum of the numbers (4 + 9 + 16)
	expectedSum := 29

	// Check if the result matches the expected sum
	if result != expectedSum {
		t.Errorf("Expected sum %d, but got %d", expectedSum, result)
	}
}

func TestSum_GracefulShutdown(t *testing.T) {
	// Create a done channel and close it immediately to simulate shutdown
	done := make(chan struct{})
	close(done)

	// Create a squareCh channel and populate it with squared numbers
	squareCh := make(chan int, 3)
	squareCh <- 4  // 2*2
	squareCh <- 9  // 3*3
	squareCh <- 16 // 4*4
	close(squareCh)

	// Call the sum function
	resultCh := sum(done, squareCh)

	// Since done is closed, the resultCh should not return any value
	select {
	case result := <-resultCh:
		expectedResult := 0
		if result != expectedResult {
			t.Errorf("Expected result %d due to partial shutdown, but got %d", expectedResult, result)
		}
	case <-time.After(10 * time.Millisecond):
		// Passed: No value should be received from resultCh
	}
}

func TestSum_PartialShutdown(t *testing.T) {
	done := make(chan struct{})
	squareCh := make(chan int)

	resultCh := sum(done, squareCh)

	// Use a separate goroutine to simulate sending values and triggering shutdown
	go func() {
		squareCh <- 4  // Send 4 (which is 2*2)
		squareCh <- 16 // Send 16 (which is 4*4)
		// Simulate a delay to ensure the value is processed before shutting down
		time.Sleep(10 * time.Millisecond)
		// Close the done channel to simulate shutdown
		close(done)
		// Attempt to send the next value, which should not be processed
		squareCh <- 9 // 3*3
		close(squareCh)
	}()

	// Since done is closed after sending two number, check the result
	select {
	case result := <-resultCh:
		// In this case, 4 and 16 was processed before shutdown, so the result should be 20
		expectedResult := 20
		if result != expectedResult {
			t.Errorf("Expected result %d due to partial shutdown, but got %d", expectedResult, result)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Expected a result, but got none")
	}
}
