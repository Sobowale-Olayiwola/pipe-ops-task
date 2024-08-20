package main

// RandIntn is a type that defines a function signature matching rand.Intn
type RandIntn func(n int) int

// randomNumberInRange generates a random number between min and max.
func randomNumberInRange(rnd RandIntn, min, max int) int {
	if min > max {
		panic("min cannot be greater than max")
	}
	return rnd(max-min+1) + min
}
