package kafka

func main() {
	// 设置Kafka客户端配置

	// 遍历所有分区，获取积压消息数量
	//for partition := range partitions {
	//	offset, err := client.GetOffset("test", int32(partition), sarama.OffsetNewest)
	//	if err != nil {
	//		log.Println("Failed to get offset for partition", partition, err)
	//		continue
	//	}
	//
	//	// 假设消费者组ID为"your_consumer_group"
	//	cgOffset, err := client.ConsumerGroupOffset("your_consumer_group", topic, int32(partition), config)
	//	if err != nil {
	//		log.Println("Failed to get consumer group offset for partition", partition, err)
	//		continue
	//	}
	//
	//	lag := offset - cgOffset
	//
	//	fmt.Printf("Partition %d has a lag of %d messages\n", partition, lag)
	//}
}
