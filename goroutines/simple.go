package main

import (
	"fmt"
	"time"
)

func worker(name string) {
	for i := 1; i <= 4; i++ {
		fmt.Println(name, "runs", i)

	}
}

// for biginners use time.sleep
func main() {
	//fmt.Println("hi from main")
	go worker("Robot 1")
	go worker("Robot 2")
	for i := 1; i < 5; i++ {

		go func() {
			fmt.Println(i)
		}()

	}
	time.Sleep(time.Second * 30)
}
