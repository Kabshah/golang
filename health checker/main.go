package main

import "fmt"

func main() {
	results := start()

	for result := range results {
		fmt.Printf(
			"[%s] %s | Status: %s | Latency: %v | Error: %s\n",
			result.Checked.Format("15:04:05"),
			result.URL,
			result.Status,
			result.Latency,
			result.Error,
		)
	}
}
