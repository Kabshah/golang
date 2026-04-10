package main

import (
	"context"
	"healthCli/model"
	"sync"
	"time"
)

// send only chan
func runChecks(url []string, outputChan chan<- model.Result) {
	var wg sync.WaitGroup
	for _, url := range url {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			outputChan <- checkUrl(u, ctx)
		}(url)
	}
	wg.Wait()
}
