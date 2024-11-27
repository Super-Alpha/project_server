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

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}

	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		return nil, err
	}

	return &Client{ClusterAdmin: admin}, nil
}

func (c *Client) CreateTopic() {
	detail := sarama.TopicDetail{
		NumPartitions:     3,
		ReplicationFactor: 2,
		ReplicaAssignment: nil,
		ConfigEntries:     nil,
	}

	if err := c.ClusterAdmin.CreateTopic("test", &detail, false); err != nil {
		log.Fatal("Error creating topic: ", err)
	}
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

func Main() {
	client, err := NewClient()
	if err != nil {
		return
	}

	defer client.ClusterAdmin.Close()

	client.DeleteRecords()
}
