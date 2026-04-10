package main

import (
	"fmt"
)

func main() {
	role := "admin"
	hasPermission := false

	if role == "guard" && hasPermission {
		fmt.Println("all in")
	} else {
		fmt.Println("Not in")
	}

	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	}

}
