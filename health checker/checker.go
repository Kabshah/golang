package main

import (
	"context"
	"healthCli/model"
	"net/http"
	"time"
)

func checkUrl(url string, ctx context.Context) model.Result {
	startTime := time.Now()
	//req, err := http.NewRequestWithContext(ctx, method, url, body)
	// 	5️⃣ Real world analogy 🍔

	// Socho tum restaurant me order dete ho.

	// Normal request:

	// burger order
	// wait forever

	// Context request:

	// burger order
	// agar 10 min me nahi aya
	// cancel order
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	// yeh line add karo 👇
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0")
	client := &http.Client{}
	resp, err := client.Do(req)
	latency := time.Since(startTime)
	if err != nil {
		return model.Result{
			URL:     url,
			Status:  "DOWN",
			Latency: latency,
			Error:   err.Error(),
			Checked: time.Now(),
		}
	}
	defer resp.Body.Close()

	status := "UP"
	if resp.StatusCode >= 400 {
		status = "DOWN"
	}
	return model.Result{
		URL:     url,
		Status:  status,
		Latency: latency,
		Checked: time.Now(),
	}
}
