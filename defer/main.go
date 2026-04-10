package main

import (
	"errors"
	"fmt"
)

func main() {
	// defer follows lifo aur aur yrr jis line sy phly defer ho it executes in last. it is mainly used for cleanups
	// lo0ped defer becomes heavy
	// guranted run at the end of function
	fmt.Println("CASE 1: SUCCESS CASE")
	doSomething(true)
	fmt.Println("CASE 2: FAILED CASE")
	doSomething(false)

}
func doSomething(sucess bool) error {
	fmt.Println("start: Resource acquired")
	defer fmt.Println("Cleanup: resource released")
	if !sucess {
		return errors.New("something went wrong iam returning.BYE!")
	}
	fmt.Println("Doing some work")
	fmt.Println("Doing something important")
	return nil
}
