package main

import (
	"fmt"
	"os/signal"

	"os"

	"github.com/Shopify/sarama"
)

func main() {

	//addresses of available kafka brokers
	brokers := []string{"rm-be-k8k73.beta.local:9092", "rm-be-k8k74.beta.local:9092", "rm-be-k8k75.beta.local:9092"}

	consumer, err := sarama.NewConsumer(brokers, nil)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	subscribe(consumer, "users-domain-api")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	<-exit
	consumer.Close()
	fmt.Println("Bye")

}

func subscribe(consumer sarama.Consumer, topic string) {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetOldest //get all the offsets
	//initialOffset := sarama.OffsetNewest // get all new initialOffset

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case err := <-pc.Errors():
					fmt.Println(err)
				case message := <-pc.Messages():
					messageReceived(message)
				}
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	fmt.Printf("%v : %s : %s : %s\n", message.Offset, message.Topic, string(message.Key), string(message.Value))
}
