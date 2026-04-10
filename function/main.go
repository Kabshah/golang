package main

import "fmt"

// multiple return types
func SumAndProduct(a int, b int) (int, int) {
	SUM := a + b
	prod := a * b
	return SUM, prod
}

// multiple parameters
func sumAll(nums ...int) int {
	total := 0
	for _, val := range nums {
		total = total + val
	}
	return total
}

// golang variadic function
//function that accepts multiple no of args of same type
func main() {
	s, p := SumAndProduct(4, 5)
	fmt.Println(s, p)
	onlySum, _ := SumAndProduct(4, 5)
	fmt.Println(onlySum)
	_, onlyProd := SumAndProduct(4, 5)
	fmt.Println(onlyProd)
	fmt.Println("multiple parameters: ", sumAll(12, 16, 18))

	// accessing multiple params through slices
	slices1 := []int{1, 2, 3}
	fmt.Println("accessing multiple params through slices", sumAll(slices1...))

	// anonymous function it doesnt have a name
	res := func(n int) int {
		return n * 2
	}
	fmt.Println("Anon func doesnt have a name :", res(2))

	// IIFE STYLE LIKE JS FUNCTION
	result := func(a int, b int) int {
		return a + b
	}(5, 5)

	fmt.Println(result)

}
