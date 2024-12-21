# Go Concurrency Bug: Data Race in Slice Appending

This repository demonstrates a subtle data race bug in a Go program that uses goroutines to concurrently append to a slice.  The program appears to use proper synchronization with a mutex, but still exhibits unexpected behavior due to the way the slice length is accessed.

## The Bug

The `bug.go` file contains the buggy code. While a `sync.RWMutex` is used to protect the append operation, there is still a data race because `len(data)` is not protected.  Each goroutine appends to the slice, but the final length calculation is done outside the mutex.  Therefore, the length might not reflect the total number of appends.

## The Solution

The `bugSolution.go` file demonstrates a corrected version of the code. By protecting the `len(data)` calculation with a read lock, the race condition is eliminated and the program consistently prints 1000.

## How to Reproduce

1. Clone this repository.
2. Navigate to the directory.
3. Run `go run bug.go` to observe the erratic behavior.
4. Run `go run bugSolution.go` to see the correct output.
