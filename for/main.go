package main

import (
	"fmt"
)

// only for-loop for looping in go nothing else
func main() {
	i := 1
	for i <= 3 {
		fmt.Print(i)
		i++
	}

	// for {
	// 	fmt.Println("infinite")
	// }

	fmt.Println("Range ")
	for i := range 20 {
		fmt.Println(i)
	}
	//  prints from 0 to 19(n-1)
}
