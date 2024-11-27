package kafka

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

/*
 * @Description: 针对kafka，采用单消费者实例订阅消息
 */

func Consumer() {
	config := sarama.NewConfig()
	// 消费者配置
	config.Version = sarama.V3_6_0_0                              // 指定kafka版本3.6.0
	config.Consumer.Return.Errors = true                          // 设置是否返回消费过程中的错误
	config.Consumer.Offsets.Initial = sarama.OffsetNewest         // 在没有提交位点情况下，使用最新的位点还是最老的位点，默认是最新的消息位点
	config.Consumer.Offsets.AutoCommit.Enable = true              // 是否支持自动提交位点，默认支持
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动提交位点时间间隔，默认1s
	config.Consumer.MaxWaitTime = 250 * time.Millisecond          // 在没有最新消费消息时候，客户端等待的时间，默认250ms
	config.Consumer.MaxProcessingTime = 100 * time.Millisecond    // 消息处理的超时时间，默认100ms，
	config.Consumer.Fetch.Min = 1                                 // 消费请求中获取的最小消息字节数,Broker将等待至少这么多字节的消息然后返回。默认值为1，不能设置0，因为0会导致在没有消息可用时消费者空转。
	config.Consumer.Fetch.Max = 0                                 // 消费请求最大的字节数。默认为0，表示不限制
	config.Consumer.Fetch.Default = 1024 * 1024                   // 消费请求的默认消息字节数（默认为1MB），需要大于实例的大部分消息，否则Broker会花费大量时间计算消费数据是否达到这个值的条件
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second    // 设置rebalance操作的超时时间，默认60s
	config.Consumer.Group.Session.Timeout = 10 * time.Second      // 设置消费者组会话的超时时间为,默认为10s
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second    // 心跳超时时间，默认为3s

	// 消费者客户端,连接到kafka集群
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}

	defer func() {
		_ = consumer.Close()
	}()

	// 订阅主题,消费指定分区中的消息,从最新位置开始消费
	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}

	defer partitionConsumer.Close()

	// 获取分区的高水位,即下一个待消费消息的位移
	// partitionConsumer.HighWaterMarkOffset()

	// 从分区中接收消息
	for {
		select {
		case msg, ok := <-partitionConsumer.Messages():
			if ok {
				log.Printf("Received message, topic = %s, partition = %d, offset = %d, key = %s, value = %s\n",
					msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}
	}
}
