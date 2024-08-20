package main

// square receives integers from genCh, squares them, and sends the results to squareCh.
func square(done <-chan struct{}, genCh <-chan int, squareCh chan<- int) {
	defer close(squareCh)

	for num := range genCh {
		select {
		case <-done: // This makes the cancellation of higher priority for immediate shutdown
			return
		default:
			select {
			case squareCh <- num * num:
			case <-done:
				return
			}
		}
	}
}
