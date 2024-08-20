package main

// sum receives squared integers from squareCh, sums them, and sends the final sum to the returned channel.
func sum(done <-chan struct{}, squareCh <-chan int) <-chan int {
	resultCh := make(chan int, 1)

	go func() {
		totalSum := 0
		defer func() {
			resultCh <- totalSum
			close(resultCh)
		}()
		for num := range squareCh {
			select {
			case <-done: // This makes the cancellation of higher priority for immediate shutdown
				return
			default:
				select {
				case <-done:
					return
				default:
					totalSum += num
				}
			}
		}
	}()
	return resultCh
}
