package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

// This function receives the values of firstName and lastName
// then returns the concatenation of both
func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return fullName
}

func main() {
	e := Employee{
		FirstName: "X Æ A-12",
		LastName:  "Musk",
	}
	// If we wanted to print the full name
	fmt.Println("Son's name was", fullName(e.FirstName, e.LastName))

	// But it sucks to send the whole thing every time
	// So we use recievers (below)
	fmt.Printf("IT WAS %s !!!\n", e.fullName())

	e.changeLastName("X AE A-Xii")
	fmt.Printf("But then they changed it to: %s\n", e.fullName())
}

// Function Receivers by value
// For quicker access to fields; attach the function to the struct itself
func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

// Function Receivers by reference
// For modifying the internals of the struct
func (e *Employee) changeLastName(firstName string) {
	e.FirstName = firstName
}
