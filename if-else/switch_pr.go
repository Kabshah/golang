package main

import (
	"fmt"
	"time"
)

func cong() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Work day")
	}
}

func main() {
	cong()

	unknown := func(t any) {
		switch i := t.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its an string")
		default:
			fmt.Println("Other: ", i)
		}
	}
	unknown(19.7)
	//unknown("kab")
}

//go run switch_pr.go
// works
