package main

import "fmt"

func changeVal(score *int) {
	*score = *score + 10
}

func main() {
	// store the memo adress of any val
	// &x -> addresS of x(makes a pointer)
	// *p -> dereferencing operator (go to that address and read/write)

	// 	x := 10
	// p := &x
	// Pointer value nahi rakhta.Pointer us room ka address rakhta hai jahan value rakhi hai
	// Pointers ka main reason hai:

	//👉 function ke andar original variable change karna
	// W/O POINTER
	// 	func change(n int) {
	// 	n = 50
	// }

	// func main() {
	// 	x := 10
	// 	change(x)
	// 	fmt.Println(x)
	// }

	// WITH POINTER
	// 	func change(n *int) {
	// 	*n = 50
	// }

	// func main() {
	// 	x := 10
	// 	change(&x)
	// 	fmt.Println(x)
	// }

	points := 10
	fmt.Println("b4 score: ", points)
	fmt.Println("Memory address", &points)
	changeVal(&points)
	fmt.Println(" After score is: ", points)
}
