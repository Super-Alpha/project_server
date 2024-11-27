package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

/*
 * @Description: 针对broker节点，执行某些操作（比如获取元数据（该节点上主题、分区、节点ID））
 */

// 连接单一节点
func Broker() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	broker := sarama.NewBroker("localhost:9092")
	err := broker.Open(config)
	if err != nil {
		panic(err)
	}

	request := sarama.MetadataRequest{Topics: []string{"test"}}
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	for _, topic := range response.Topics {
		for _, partition := range topic.Partitions {
			fmt.Printf("Topic: %s, Partition: %d, Leader: %d\n", topic.Name, partition.ID, partition.Leader)
		}
	}

	if err = broker.Close(); err != nil {
		panic(err)
	}
}

// 连接多个broker节点
func Brokers() {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	// 三个节点broker
	brokersList := []string{
		"localhost:9092",
		"localhost:9093",
		"localhost:9094",
	}

	client, err := sarama.NewClient(brokersList, config)
	if err != nil {
		log.Fatalln("Failed to initialize Kafka client:", err)
	}

	defer client.Close()

	brokers := client.Brokers()

	for _, broker := range brokers {

		if err := broker.Open(config); err != nil {
			log.Fatalln("Failed to connect to broker:", err)
		}

		ok, err := broker.Connected()
		if err != nil {
			log.Fatalln("Failed to get broker status:", err)
		}

		if ok {
			metaReq := &sarama.MetadataRequest{
				Topics:                             []string{"test"},
				IncludeClusterAuthorizedOperations: true,
				IncludeTopicAuthorizedOperations:   true,
			}

			resp, err := broker.GetMetadata(metaReq)
			if err != nil {
				log.Fatalln("Failed to get metadata:", err)
			}

			fmt.Printf("Broker ID: %d, Address: %s, ClusterID: %d\n", broker.ID(), broker.Addr(), resp.ControllerID)
		}

		if err = broker.Close(); err != nil {
			panic(err)
		}
	}
}
