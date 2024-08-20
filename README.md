# Go Pipeline Operations Take-Home Test

## Overview

This one-day take-home test is designed to assess your practical understanding of pipeline operations in Go. You'll be implementing a simple data processing pipeline that demonstrates your ability to work with goroutines and channels.

## Task Description

Create a Go program that processes a stream of integers using a pipeline approach. The program should consist of the following stages:

1. **Generator**: Create a function that generates a stream of random integers (between 1 and 100) and sends them to a channel.

2. **Square**: Implement a stage that receives numbers from the generator, squares them, and sends the results to the next stage.

3. **Sum**: Implement a final stage that sums all the numbers it receives and prints the final result.

## Requirements

1. Use goroutines to run each stage of the pipeline concurrently.
2. The generator should produce 10,000 numbers and above.
3. Use buffered channels where appropriate.
4. Include comments explaining your code and design decisions.
5. Write at least one unit test for each stage of the pipeline.
6. Implement proper error handling and graceful shutdown of the pipeline.

## Bonus Points (Optional)

- Add a simple command-line flag to control the number of numbers generated.
- Print out how long the entire pipeline takes to process all numbers.

## Submission Guidelines

1. Create a GitHub repository for your solution.
2. Include a `main.go` file with your pipeline implementation.
3. Write unit tests in a `main_test.go` file.
4. Update this README with instructions on how to run your program and tests.

## Evaluation Criteria

Your solution will be evaluated based on:

1. Correctness of the pipeline implementation
2. Proper use of Go concurrency primitives (goroutines, channels)
3. Code organization and clarity
4. Error handling
5. Quality of unit tests

## Time Allocation

This test is designed to be completed within one day (about 4-6 hours of focused work). Please submit your solution within 24/48 hours of receiving the test.

Good luck! We look forward to reviewing your solution.

## How To Run And Generate Executable

To generate a build and file run with command line flag (**n**) where **n** should be equal or more than 10,000, if **n** is less than, the application defaults to 10,000

* go build

* ./pipeline -n=15000

## Test And Coverage

To generate test report run
* go test -v

To generate test coverage
* go test -cover
