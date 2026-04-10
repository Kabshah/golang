package order

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consumer() {
	// kafka consumer
	// for reading input from producer
	reader := kafka.NewReader(kafka.ReaderConfig{
		// there can be multiple brokers running on multiple server from architectural pov
		Brokers: []string{"localhost:29092"},
		// topic subscribed by this consumer is orders
		// this consumer is subscribed to orders topic
		Topic: "orders",
		// watchers, dressers,
		GroupID: "orders-group",
	})
	defer reader.Close()
	fmt.Println("Consumer listening...")
	// for each msg sent from producer we are going to read
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var o Order // is is var implementing order struct usko hum read hoi wi val dey rahi for validation purpose
		json.Unmarshal(msg.Value, &o)
		// r%s Product name %d(value int)
		fmt.Printf("Order received %s:%d\n", o.Product, o.Quantity)
	}

}
