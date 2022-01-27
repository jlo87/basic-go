package main

import (
	"fmt"
)

// Go is a functional language, which means that functions are first-class citizens
func main() {

	// defer postpones the implications of a function
	// the order of invoking defer statements are reversed
	defer LastHi() // this is invoked last
	defer func() { // second defer will be invoked second to last
		fmt.Println("Almost Last Hi")
	}()

	// functions can also be values
	a := Hello()
	fmt.Println(a)

	// Can return multiple values
	b, c := TwoValues()
	fmt.Println(b, c)

	// Functions are values
	d := TwoValues
	fmt.Println(d())
}

// this function returns a single value
func Hello() string {
	return "Hello There"
}

// this function returns two values(both strings)
func TwoValues() (string, string) {
	return "Hello", "World"
}

func LastHi() {
	fmt.Println("Last Hi!")
}
