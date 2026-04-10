package main

import (
	"fmt"
	"sync"
)

// 1) ADD 2) DONE 3) WAIT for wg

func task(id int, wg *sync.WaitGroup) {
	fmt.Println("Doing task", id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
