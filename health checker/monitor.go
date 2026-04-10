package main

import (
	"healthCli/model"
	"time"
)

var services = []string{
	"https://google.com",
	"https://github.com",
	"https://stackoverflow.com",
}

func start() chan model.Result {
	outputChan := make(chan model.Result)
	go func() {
		ticker := time.NewTicker(time.Second * 5)

		//infinite for loop
		for {
			runChecks(services, outputChan)
			<-ticker.C
		}
	}()
	return outputChan
}
