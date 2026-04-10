package main

import "fmt"

// unbuffered = ek time ek hee val/data can be send jb tk wo receive nai hota h dosra send nai krr skty h
// buffered= limited amount of data can be send w/o blocking
func main() {
	// is limt k baad ka jo data hga na wo blocking rahaiga
	// first time in channels basic ka commented code deadlock error aya tha but yahn deko nai aya h

	emailSend := make(chan string, 100)
	// agr iss capacity sy zyda send kara toh blocking hoga aur deadlock error ayega
	//emailSend := make(chan string, 2)
	emailSend <- "iam@gmail.com"
	emailSend <- "iam1@gmail.com"
	// emailSend <- "iam2@gmail.com"
	fmt.Println(<-emailSend)
	fmt.Println(<-emailSend)
	//fmt.Println(<-emailSend)
}
