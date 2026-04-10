package main

import (
	"kafka-go/internal/order"
	"log"
	"net/http"
)

func main() {
	order.CreateTopic()
	go order.Consumer()
	http.HandleFunc("/", order.OrderHandler)
	http.HandleFunc("/order", order.OrderHandler)
	log.Println("Starting server at the localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
