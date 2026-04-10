package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	// anonymous func
	go func() {
		chan1 <- 19
	}()
	go func() {
		chan2 <- "kabshah"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan_int := <-chan1:
			fmt.Println("its int", chan_int)

		case chan_string := <-chan2:
			fmt.Println("its a string,", chan_string)
		}
	}
}
