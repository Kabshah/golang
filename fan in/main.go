package main

// fan in workers agr data return krr rhy h uss data ko huma agregrate krna hota h
import (
	"fmt"
	"sync"
	"time"
)

type ImgInfo struct {
	urlVal string
	Err    error
}

func worker(url string, wg *sync.WaitGroup, resultChan chan ImgInfo) {
	defer wg.Done()
	fmt.Printf("Image processed: %s\n", url)

	//sending our url to chanel
	resultChan <- ImgInfo{
		urlVal: url,
		Err:    nil,
	}

}
func main() {
	var wg sync.WaitGroup
	startTime := time.Now()
	resultChan := make(chan ImgInfo, 2)

	wg.Add(2)

	go worker(`C:\Users\User\Pictures\Screenshots\New folder`, &wg, resultChan) //go routine
	go worker(`C:\Users\User\Pictures\Screenshots\New folder`, &wg, resultChan) //go routine
	// channels go routine sa data main maa laa rahy like this

	//  Alag goroutine: jab sab workers done ho jayen, channel band karo
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Printf("Received: %s\n", result)
	}

	fmt.Printf("It took %v\n", time.Since(startTime))

}

// yhn sirf humari jobs process hokr chan sy ari h
