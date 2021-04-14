package main

// Imports are done this way
import "fmt"

func main() {
	// For is everything

	// Simple For loop
	sum := 0
	for i := 1; i < 5; i++ { // Variable i; stopping condition; init i (i incremented)
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)

	// Simple While loop (for is Go's while)
	n := 1
	for n < 5 { // n < 5 is the stopping condition
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	// Infinite loop
	// sum := 0
	// for {      // Nothing provided to for
	// 		sum++ // repeated forever/iterate indefinitely
	// stop loop by using break; or return to get out of function
	// }
	// fmt.Println(sum) // never reached
}
