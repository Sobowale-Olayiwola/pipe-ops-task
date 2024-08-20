package main

// generator generates a stream of random integers and sends them to the genCh channel.
func generator(rnd RandIntn, done <-chan struct{}, n int, genCh chan<- int) {
	defer close(genCh)

	for i := 0; i < n; i++ {
		select {
		case <-done: // This makes the cancellation of higher priority for immediate shutdown
			return
		default:
			select {
			case genCh <- randomNumberInRange(rnd, minRandNum, maxRandNum):
			case <-done:
				return
			}
		}
	}
}
