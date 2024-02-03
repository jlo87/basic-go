/*
Concurrency allows us to have multiple blocks of code share the
execution time by pausing their execution normally. We can also
run blocks of code in parallel at the same time. Concurrent tasks
in Go are called go routines, and not threads. To execute multiple
functions in new go routines, all we need to do is add the word go
in front of the function.
*/

package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Fun 1 :", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Fun 2 :", i)
	}
}

func main() {
	// When using go routine, you cannot trust in what order
	// they are going to execute - so we need channels
	go printTo15()
	go printTo10()

	// We need to pause the execution of the main function
	// for other functions to execute
	time.Sleep(2 * time.Second)
}
