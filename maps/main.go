package main

import "fmt"

// Maps specifies the key and value. Key [string] enclosed in square brackets
// and return value is the int.
var sampleMap map[string]int

// A complex map of string that returns a map of string that
// returns a string
var complexSampleMap map[string]map[string]string

func main() {
	sampleMap = map[string]int{
		"James": 50,
		"Ali":   39,
	}

	// Currency map. Key is string and return value
	// is also a string
	currency := map[string]string{
		"AUD": "Australia Dollar",
		"GBP": "Great Britain Pound",
		"JPY": "Japan Yen",
		"CHF": "Switzerland Franc",
	}

	// a. Adding to the map:
	// Modify elements within a map by accessing the key
	currency["USD"] = "USA Dollar"

	fmt.Println("Currency with USD added: ", currency)

	// b. Remove from the map:
	// Delete members of the map
	delete(currency, "GBP")
	fmt.Println("Currency with GBP deleted: ", currency)

	// c. Replacing one entry with another:
	currency["AUD"] = "New Zealand Dollar"
	fmt.Println("Currency with AUD value replaced with NZD: ", currency)

	// Ranging through the map:
	for key, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", key, value)
	}

	// If you don't want to use both key and value, ignore the
	// key by using a placeholder
	for _, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", value)
	}

	// If only taking key, you don't need the value
	for key := range currency {
		fmt.Printf("%v might be equal to: %v\n", key)
	}
}
