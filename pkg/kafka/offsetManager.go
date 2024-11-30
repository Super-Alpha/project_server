package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

func OffsetManager() {

	config := sarama.NewConfig()

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	offsetManagerClient, err := sarama.NewOffsetManagerFromClient("test-group", client)
	if err != nil {
		panic(err)
	}

	defer offsetManagerClient.Close()

	partitionOffsetManager, err := offsetManagerClient.ManagePartition("test", 0)
	if err != nil {
		panic(err)
	}

	nextOffset, res := partitionOffsetManager.NextOffset()

	fmt.Printf("NextOffset %d,res %s\n", nextOffset, res)
}
