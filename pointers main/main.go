package main

import "fmt"

func change_by_val(num *int) {
	*num = 5 // we r changing to this num
	//fmt.Println()
}
func main() {
	// jb humko main ka value change krna hota h dosry func ma sy toh we use pointers
	// normal agr func sy change kary w/o pointers toh wo change by reference khlata h

	num := 1
	change_by_val(&num)
	fmt.Println("changed val is: ", num)

}
