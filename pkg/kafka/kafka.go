package kafka

import "github.com/IBM/sarama"

func main() {
	// 设置Kafka客户端配置
	config := sarama.NewConfig()

	client, err := sarama.NewClient([]string{}, config)
	if err != nil {
		panic(err)
	}

	defer client.Close()
}
