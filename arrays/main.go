package main

import "fmt"

func main() {
	// we use slices mainly in go
	fixed := [4]int{1, 2, 3, 4} //arr literal khty hein issy fixed walay hko
	fmt.Println(fixed)

	ai_models := [3]string{"mid journey", "gemini", "deepseek"}
	fmt.Println(ai_models)

	var arr [5]int
	fmt.Println(arr)
	// output: [0 0 0 0 0]
	fmt.Println(len(arr)) // 5

	var vals [4]bool
	fmt.Println(vals)
	// output: [false false false false]

	arr2d := [2][2]int{{1, 3}, {2, 4}}
	fmt.Println(arr2d)
	// output: [[1 3] [2 4]]
}
