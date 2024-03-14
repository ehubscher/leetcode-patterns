package main

import (
	"fmt"
	"time"
)

// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
//   defer timer("sum")()
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	ContainsDuplicate("ContainsDuplicate()", nums)
}