package order

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Producer(order Order) {
	// converting model to josn format
	data, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		//ARR OF string
		Brokers: []string{"localhost:29092"},
		Topic:   "orders",
	})
	defer writer.Close()
	time.Sleep(1 * time.Second)

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key: []byte(order.Product),
		// since its already marshel no need to convert to arr of bytes
		Value: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Producer sent a message")
}
