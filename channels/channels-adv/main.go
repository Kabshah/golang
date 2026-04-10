package main

import "fmt"

// receiveing from  channel to go routine
func processNum(result chan int, num1 int, num2 int) {
	logic := num1 + num2
	result <- logic //sending logic to go routine
}
func main() {
	result := make(chan int)
	go processNum(result, 4, 5)
	receiving := <-result //receing from chan to go routine(blocking h thats why no need to sleep)
	fmt.Println(receiving)
}
