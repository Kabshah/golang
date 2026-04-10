package order

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// we are going to create a topic order which is going to be subscribed by our producer and consumer for deleivering the order

func CreateTopic() {
	conn, err := kafka.Dial("tcp", "localhost:29092")
	if err != nil {
		log.Fatal(err)
	}
	//leader broker can create topics only
	// getting leader broker of the particular kafka cluster
	// getting controller broker
	controller, err := conn.Controller()
	if err != nil {
		log.Fatal(err)
	}
	// getting host name and port no(d)
	controllerConn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer controllerConn.Close()

	topicName := "orders"

	err = controllerConn.CreateTopics(kafka.TopicConfig{
		Topic:             topicName,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Created topic Orders sucessfully")
	}

}
