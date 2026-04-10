package main

import "fmt"

// Universal formatter hai.
// Almost kisi bhi type ko print kar deta hai (string, int, struct, slice, map, etc.).

func main() {
	name := "kabshah"
	age := 19

	fmt.Printf("Name: %v Age: %v\n", name, age)

	// %s Sirf string print karne ke liye use hota hai.
	// Agar variable string type ka hai to usay direct print karta hai.

	name2 := "Aaima"
	fmt.Printf("Name: %s\n", name2)
}
