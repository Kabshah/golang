package main

import "fmt"

func main() {
	// common collection type
	// dynamic and it can grow

	results := []string{"sangam", "kabshah", "john"} //slice
	fmt.Println(results[0], "last one is: ", results[len(results)-1])

	//slice operator
	var num5 = []int{1, 2, 3}
	fmt.Println((num5[0]))
	fmt.Println(num5[1])
	fmt.Println(num5[2])
	fmt.Println(num5[0:2])

	//another way of writing a slice which will grow
	var nums []int
	nums = append(nums, 10)
	nums = append(nums, 15, 16)
	fmt.Println(nums)
	// capacity how many u can store
	// it creates a slice with a given length an capacity

	// make([]t, len,cap)
	scores := make([]int, 0, 5)
	fmt.Println(scores, len(scores), cap(scores))

	// if u exceed the capacity go will grow it
	scores = append(scores, 100, 200, 400, 500, 600, 700)
	fmt.Println(scores)
	fmt.Println("capcity after increasing limit: ", cap(scores))

	todos := []string{" DO YOUTUBE", "workout ever"}
	more := []string{"learn golang"}

	// spread operator laga krr 2 slices ko apas ma jor dya h
	todos = append(todos, more...)
	fmt.Println(todos)

	// for range
	marks := []int{34, 89, 93, 100, 34, 96, 47}
	total := 0
	for i, marks := range marks {
		fmt.Println("index no: ", i, "students grades", marks)
		total = total + marks
	}
	fmt.Println(total)

	num3 := []int{}
	fmt.Println(num3)                       //[]
	fmt.Println("capacity is: ", cap(num3)) //output: 0

	//2d slices
	var _2dslice = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(_2dslice)

}
