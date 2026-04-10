package main

import (
	"fmt"
)

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"kabshah": 20,
	}
	fmt.Println(ages["kabshah"])

	// make(map[K]V)
	users := map[string]string{
		"u1": "dusted",
		"u2": "rusted",
		"u3": "fucked",
	}
	fmt.Println(users)
	delete(users, "u3")
	fmt.Println(users)

	// read val ok (for safe checks we use this)
	POINTS := map[string]int{
		"a": 10,
		"b": 15, //it is valid
	}
	fmt.Println(POINTS["b"])
	fmt.Println(" c value is:", POINTS["c"]) // c value is: 0

	valB, okB := POINTS["b"]
	fmt.Println(valB, okB) //output 10 True

	if val, ok := POINTS["c"]; ok {
		fmt.Println(val, "b present")
	} else {
		fmt.Println("B aint present")
	}
	prices := map[string]int{
		"guy1": 1800,
		"guy2": 1400,
	}
	//total := 0
	for u := range prices {
		fmt.Println(u)
	}
}
