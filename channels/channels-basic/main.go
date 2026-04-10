package main

import (
	"fmt"
	"time"
)

// passage h b/w goroutines for sending and receiving data

// sending from goroutine to channel
func processStr(messageChan chan string) {
	fmt.Println("Processing", <-messageChan)
}
func main() {
	// basic and easy saa channel
	messageChan := make(chan string) //unbuffered channel
	//messageChan <- "ping" // for sending data (channels are blocking)
	// msg := <-messageChan  // for receiving chan
	// fmt.Println(msg)
	go processStr(messageChan)
	messageChan <- "ping" //phly go routine start karo then send karo takay go routine receive kary

	time.Sleep(time.Second * 2)
}
