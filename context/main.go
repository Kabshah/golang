package main

import (
	"context"
	"fmt"
	"time"
)

//timeout code (long running gorroutines can be cancelled with context)
// sharing data in code
//db op can rely on context for cancellation

func longRunningfunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("CONTEXT CANCELLING...")
			return
		default:
			fmt.Println("Doing some work")
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel() //last cleaning func yeh print krrwaa rha h
	go longRunningfunc(ctx)
	time.Sleep(time.Second * 5) //this is for goroutine
	cancel()                    //actually cancelling and stoping "doing some work"
}
