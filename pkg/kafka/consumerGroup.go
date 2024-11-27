package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

/*
 * @Description: 针对kafka，采用消费者组订阅消息
 */

func ConsumerGroup() {

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 针对消费者组的分区分配策略
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.ChannelBufferSize = 1024 // channel长度

	var wg sync.WaitGroup

	wg.Add(2)

	// 同一个消费者组“test-group”，开启两个消费者
	go func() {
		defer wg.Done()
		if err := CreateConsumerAndConsume(context.Background(), config, "consumer1"); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := CreateConsumerAndConsume(context.Background(), config, "consumer2"); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}

func CreateConsumerAndConsume(ctx context.Context, config *sarama.Config, consumerName string) error {

	// 创建client
	newClient, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		return err
	}

	// 获取所有的topic
	tps, err := newClient.Topics()
	if err != nil {
		return err
	}
	fmt.Printf("topics: %s\n", tps)

	// 根据client创建consumerGroup客户端
	client, err := sarama.NewConsumerGroupFromClient("test-group", newClient)
	if err != nil {
		return err
	}

	defer func() {
		_ = client.Close()
	}()

	handlerObject := &ConsumerGroupHandler{
		ConsumerName: consumerName,
	}

	for {
		// 该方法要在循环中调用，否则会阻塞
		if err = client.Consume(ctx, []string{"test"}, handlerObject); err != nil {
			return err
		}
	}
}

type ConsumerGroupHandler struct {
	ConsumerName string
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("Setup consumerName = %s, Partition = %+v\n", c.ConsumerName, session.Claims())
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("cleanup")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("consumerName = %s, memberId:%s, topic:%s, partiton:%d, offset:%d, value:%s, timeStamp:%v\n",
			c.ConsumerName, session.MemberID(), message.Topic, message.Partition, message.Offset, string(message.Value), message.Timestamp.Format("2006-01-02 15:04:05"))
		// 更新位移,标记消息已经被消费
		session.MarkMessage(message, "consumed") // 自动提交位移
		// session.Commit() 手动提交位移
	}
	return nil
}
