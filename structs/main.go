package main

import (
	"fmt"
	"sync"
)

// Structs are very similar to other objects in OOP languages
// A struct is basically a class in other languages; a type in Go
// This is called a Person Type
type Person struct {
	FirstName string `json: "firstName" bson: "firstName" yaml: "firstName"` // Tagging can be added for serialization
	LastName  string `json: "lastName" bson: "lastName" yaml: "lastName"`    // Tagging is just added metadata on each property of the struct
	Age       int
}

func main() {
	// Initialize a person, of type Person
	// Assigning values to the fields in person struct
	p1 := Person{
		FirstName: "Jonathan",
		LastName:  "Lo",
		Age:       34,
	}

	fmt.Println("His name is: ", p1)

	// Animal struct defined inside the main function
	type Animal struct {
		Name            string
		Characteristics []string
	}

	// Initialize the animal struct
	animal1 := Animal{
		Name: "Lion",
		Characteristics: []string{"Eats other animals",
			"Wild Animal",
			"King of the jungle",
		},
	}

	// We use the dot(.) to access each field in the struct
	// Iterate over the slice of Animal
	fmt.Println("Animal name: ", animal1.Name)
	for _, v := range animal1.Characteristics {
		fmt.Printf("\t %v\n", v)
	}

	// Nested Structs

	// An Herbivore is an animal, so it can have the animal struct as a field
	// Promotion - allows for complex types inside of structs
	type Herbivore struct {
		Animal     Animal
		EatAnimals bool
		sync.Mutex // State that this type also has a Mutex
	}

	herb := Herbivore{
		Animal: Animal{
			Name: "Goat",
			Characteristics: []string{"Lacks sense",
				"Lazy",
				"Eats grass",
			},
		},
		EatAnimals: false,
	}

	// Lock this object; Mutex inside that we are locking
	herb.Lock()
}
