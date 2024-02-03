package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Alias for fmt.Println
var pl = fmt.Println

// func main will get the user input
func main() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)

	// Read up until the point in which we get a new line,
	// or the user inside of the terminal hits the enter key
	name, err := reader.ReadString('\n')

	// Handle error
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}
