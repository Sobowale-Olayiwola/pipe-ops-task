package main

import (
	"flag"
	"fmt"
	"time"

	"math/rand"
)

const (
	minRandNum = 1
	maxRandNum = 100
)

func runPipeline(rnd RandIntn, numCount int) int {
	start := time.Now()

	// Channels

	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	genCh := make(chan int, 100)
	squareCh := make(chan int, 100)

	// Pipeline stages
	go generator(rnd, done, numCount, genCh)
	go square(done, genCh, squareCh)
	finalSum := sum(done, squareCh)

	fmt.Printf("Final sum: %d\n", finalSum)
	fmt.Printf("Time taken for pipeline completion: %v\n", time.Since(start))

	return <-finalSum
}

func main() {
	numCount := flag.Int("n", 10_000, "Number of random integers to generate")
	flag.Parse()

	if *numCount < 10_000 {
		*numCount = 10_000
	}

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rng := rand.New(source)
	runPipeline(rng.Intn, *numCount)
}
