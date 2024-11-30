package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

/*
 * @Description: ClusterAdmin是Kafka的管理客户端，支持管理和检查topics、brokers、configurations和访问控制列表（ACLs）。
 */

type Client struct {
	ClusterAdmin sarama.ClusterAdmin
}

func NewClient() (*Client, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	client, err := sarama.NewClient([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, config)
	if err != nil {
		return nil, err
	}

	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		return nil, err
	}

	return &Client{ClusterAdmin: admin}, nil
}

func (c *Client) CreateTopic(topic string) error {
	detail := sarama.TopicDetail{
		NumPartitions:     3, // 三个分区
		ReplicationFactor: 2, // 每个分区两个副本
		ReplicaAssignment: nil,
		ConfigEntries:     nil,
	}

	if err := c.ClusterAdmin.CreateTopic(topic, &detail, false); err != nil {
		return err
	}

	return nil
}

func (c *Client) ConsumerGroups() {
	consumerGroup, err := c.ClusterAdmin.ListConsumerGroups()
	if err != nil {
		log.Fatal("Error listing consumer group: ", err)
	}

	for group, _ := range consumerGroup {
		log.Println(group)
	}
}

func (c *Client) DeleteRecords() {
	// 删除主题“test”中分区1中，offset < 5的消息
	if err := c.ClusterAdmin.DeleteRecords("test", map[int32]int64{1: 5}); err != nil {
		log.Fatal("Error deleting records: ", err)
	}
}

func (c *Client) DeleteTopic(topic string) {
	if err := c.ClusterAdmin.DeleteTopic(topic); err != nil {
		log.Fatal("Error deleting topic: ", err)
	}
}

func ClusterAdmin() {
	client, err := NewClient()
	if err != nil {
		return
	}

	defer client.ClusterAdmin.Close()

	if err = client.CreateTopic("demo"); err != nil {
		log.Fatal("Error creating topic: ", err)
	}
}
