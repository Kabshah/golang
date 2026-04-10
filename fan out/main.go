// multiple go routines banao aur unko kaam dey dou distribute karo = fan out
// fan in workers agr data return krr rhy h uss data ko huma agregrate krna hota h

package main

import (
	"fmt"
	"time"
)

func worker(url string) {
	fmt.Println("Processed Image: ", url)
}
func main() {
	go worker(`C:\Users\User\Pictures\Screenshots\New folder`)
	go worker(`C:\Users\User\Pictures\Screenshots\New folder`)
	go worker(`C:\Users\User\Pictures\Screenshots\New folder`)

	time.Sleep(time.Second * 3)

	startTime := time.Now()
	elapsed := time.Since(startTime)
	fmt.Printf("It took %d ms.\n", elapsed.Milliseconds())
}
