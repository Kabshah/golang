package main

import (
	"fmt"
	"sync"
	"time"
)

// jitna jobs agr utna hee worker chala dya toh perfromance system ka degrade hojai ga
// ab bss 5 worker rkhy h saray jobs ek jagah jaa rahi then hum wahan sy uthaa rhy!!
// data db sy askta kahi sy b
// 2nd chan is for main goroutine -> to worker
type ImgInfo struct {
	urlVal string
	Err    error
}

func worker(jobsChan chan string, resultChan chan ImgInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	// receiving jobs for processing
	for job := range jobsChan {
		resultChan <- ImgInfo{
			urlVal: job,
			Err:    nil,
		}
		fmt.Println("Image processed ", job)
	}
}
func main() {
	var wg sync.WaitGroup
	jobs := []string{
		`C:\Users\User\Pictures\Screenshots\New folder`,
		`C:\Users\User\Pictures\Screenshots\New folder`,
	}
	startTime := time.Now()
	jobsChan := make(chan string, len(jobs))
	resultChan := make(chan ImgInfo, 50)

	fmt.Println("Length of jobs is:", len(jobs))

	workers := 5
	for i := 0; i <= workers; i++ {
		wg.Add(1)
		go worker(jobsChan, resultChan, &wg)
	}

	//iterator 0 sy chl rha h aur jobs hein 1,2 islia use < only not'='
	for x := 0; x < len(jobs); x++ {
		jobsChan <- jobs[x]
	}
	close(jobsChan) // forever waiting naa ho islia yhn hee band krdya

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Printf("Received: %s\n", result)
	}
	fmt.Printf("It takes: %v\n", time.Since(startTime))
}

// sending the jobs to worker [1st chan]
// processing jobs and returning them *resultChan
