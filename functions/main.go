package main

import (
	"fmt"
)

func main() {
	defer LastHi()
	defer func() {
		fmt.Println("Almost Last Hi")
	}()

	a := Hello()
	fmt.Println(a)

	// Can return multiple values
	b, c := TwoValues()
	fmt.Println(b, c)

	// Functions are values
	d := TwoValues
	fmt.Println(d())
}

func Hello() string {
	return "Hello There"
}

func TwoValues() (string, string) {
	return "Hello", "World"
}

func LastHi() {
	fmt.Println("Last Hi!")
}
