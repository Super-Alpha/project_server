package kafka

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

/*
 * @Description: 生产者发送消息
 */

func Producer() {

	config := sarama.NewConfig()
	//生产者配置
	config.Producer.Partitioner = sarama.NewHashPartitioner            // 使用Hash分区器,按照hash值来分配到不同的分区中
	config.Producer.RequiredAcks = sarama.WaitForAll                   // 设置生产者需要等待所有分区确认消息写入成功
	config.Producer.Idempotent = true                                  // 设置使用幂等性校验，确保只会有一个消费被写入
	config.Producer.Retry.Max = 3                                      // 生产者重试的最大次数，默认为3
	config.Producer.Retry.Backoff = 100 * time.Millisecond             // 生产者重试之间的等待时间，默认为100毫秒
	config.Producer.Return.Successes = false                           // 是否返回成功的消息，默认为false
	config.Producer.Return.Errors = true                               // 是否返回失败的消息，默认值为true
	config.Producer.Compression = sarama.CompressionNone               // 是否返回失败的消息，默认值为true
	config.Producer.CompressionLevel = sarama.CompressionLevelDefault  // 指定压缩等级，在配置了压缩算法后生效
	config.Producer.Flush.Frequency = 0                                // producer缓存消息的时间, 默认缓存0毫秒，即多长时间发送一批次消息
	config.Producer.Flush.Bytes = 0                                    // 达到多少字节时，触发一次发送请求，默认为0，直接发送，存在天然上限值MaxRequestSize，因此默认最大100MB，即批次达到多大，就发送一次消息
	config.Producer.Flush.Messages = 0                                 // 达到多少条消息时，强制，触发一次broker请求，这个是上限值，MaxMessages < Messages
	config.Producer.Flush.MaxMessages = 0                              // 最大缓存多少消息，默认为0，有消息立刻发送，MaxMessages设置大于0时，必须设置 Messages,且需要保证：MaxMessages > Messages
	config.Producer.Timeout = 5 * time.Second                          // 发送超时时间，默认5秒
	config.Producer.Transaction.Timeout = 1 * time.Minute              // 事务超时时间默认1分钟
	config.Producer.Transaction.Retry.Max = 50                         // 事务重试时间
	config.Producer.Transaction.Retry.Backoff = 100 * time.Millisecond // 事务重试间隔
	config.Producer.Transaction.ID = "test"                            // 事务ID

	// 生产者客户端，连接到Kafka集群
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	defer func() {
		_ = producer.Close()
	}()

	msgList := []map[string]string{
		{"p0": "hello partition0-1"},
		{"p0": "hello partition0-2"},

		{"p1": "hello partition1-1"},
		{"p1": "hello partition1-2"},

		{"p2": "hello partition2-1"},
		{"p2": "hello partition2-2"},
	}

	for _, msg := range msgList {
		for k, v := range msg {
			message := &sarama.ProducerMessage{
				Topic:     "test",                  // 主题
				Key:       sarama.StringEncoder(k), // 消息键,相同的键值，则发送到同一个分区中
				Value:     sarama.StringEncoder(v), // 消息内容
				Partition: 0,                       // 发送到指定分区
			}

			partition, offset, err := producer.SendMessage(message)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			fmt.Printf("Message is stored in partition %d with offset %d\n", partition, offset)
		}
	}
}
