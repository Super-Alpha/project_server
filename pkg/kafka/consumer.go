package kafka

import (
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

/*
 * @Description: 针对kafka，采用单消费者实例订阅消息
 */

func Consumer() {
	config := sarama.NewConfig()
	// 消费者配置
	config.Consumer.Return.Errors = true                          // 设置是否返回消费过程中的错误
	config.Consumer.Offsets.Initial = sarama.OffsetNewest         // 在没有提交位点情况下，使用最新的位点还是最老的位点，默认是最新的消息位点
	config.Consumer.Offsets.AutoCommit.Enable = true              // 是否支持自动提交位点，默认支持
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动提交位点时间间隔，默认1s
	//config.Consumer.MaxWaitTime = 250 * time.Millisecond          // 在没有最新消费消息时候，客户端等待的时间，默认250ms
	//config.Consumer.MaxProcessingTime = 100 * time.Millisecond    // 消息处理的超时时间，默认100ms，
	//config.Consumer.Fetch.Min = 1                                 // 消费请求中获取的最小消息字节数,Broker将等待至少这么多字节的消息然后返回。默认值为1，不能设置0，因为0会导致在没有消息可用时消费者空转。
	//config.Consumer.Fetch.Max = 0                                 // 消费请求最大的字节数。默认为0，表示不限制
	//config.Consumer.Fetch.Default = 1024 * 1024                   // 消费请求的默认消息字节数（默认为1MB），需要大于实例的大部分消息，否则Broker会花费大量时间计算消费数据是否达到这个值的条件
	//config.Consumer.Group.Rebalance.Timeout = 60 * time.Second    // 设置rebalance操作的超时时间，默认60s
	//config.Consumer.Group.Session.Timeout = 10 * time.Second      // 设置消费者组会话的超时时间为,默认为10s
	//config.Consumer.Group.Heartbeat.Interval = 3 * time.Second    // 心跳超时时间，默认为3s

	// 消费者客户端,连接到kafka集群
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}

	partitionConsumer, err := consumer.ConsumePartition("demo", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}

	defer func() {
		_ = consumer.Close()
		_ = partitionConsumer.Close()
	}()

	// 从分区中接收消息
	for {
		select {
		case msg, ok := <-partitionConsumer.Messages():
			if ok {
				log.Printf("Received message, topic = %s, partition = %d, offset = %d, key = %s, value = %s\n",
					msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		case err = <-partitionConsumer.Errors():
			log.Fatal(err)
		}
	}
}

// GetAllMessagesByTopic 获取指定topic的所有消息
func GetAllMessagesByTopic(topic string) (message map[int32][]string, err error) {

	config := sarama.NewConfig()

	client, err := sarama.NewClient([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, config)
	if err != nil {
		return
	}

	defer func() {
		_ = client.Close()
	}()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return
	}

	defer func() {
		_ = consumer.Close()
	}()

	partitions, err := client.Partitions(topic)
	if err != nil {
		return
	}

	var wg sync.WaitGroup

	wg.Add(len(partitions))

	message = make(map[int32][]string)

	for _, partition := range partitions {

		consumePartition, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Fatal(err)
		}

		go func(partitionConsumer sarama.PartitionConsumer, partition int32) {

			defer wg.Done()

			for {
				select {
				case msg, ok := <-partitionConsumer.Messages():
					if ok {
						message[partition] = append(message[partition], string(msg.Value))
						//fmt.Printf("Received message, topic = %s, partition = %d, offset = %d, key = %s, value = %s\n", topic, partition, msg.Offset, msg.Key, msg.Value)
					}

				case err, ok := <-partitionConsumer.Errors():
					if ok {
						log.Printf("Error on partition %d: %v\n", partitions[0], err)
					}

				case <-time.After(2 * time.Second):
					_ = partitionConsumer.Close()
					log.Printf("partition %d has closed\n", partition)
					return
				}
			}
		}(consumePartition, partition)
	}

	wg.Wait()

	return
}
