package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

/*
 * @Description: ACL(Access Control Lists)用于管理Kafka的访问控制权限。
		通过ACL,可以为Kafka中的资源(如topic、consumerGroup、broker等)设置权限，控制哪些客户端可以访问这些资源，以及它们可以执行哪些操作。
		ACL是Kafka提供的用于增强安全性的机制，主要用于限制不授权的客户端访问Kafka集群。
*/

func mains() {
	// 创建 AdminClient
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka client: %v", err)
	}

	defer client.Close()

	adminClient, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		log.Fatalf("Failed to create Admin client: %v", err)
	}

	defer adminClient.Close()

	// 创建 ACL
	acl := &sarama.ResourceAcls{}
	acl.Resource = sarama.Resource{
		ResourceType: sarama.AclResourceTopic,
		ResourceName: "test-topic",
	}
	acl.Acls = []*sarama.Acl{
		{
			Principal: "User:TestUser",
			Host:      "*",
			Operation: sarama.AclOperationRead,
		},
	}

	// 创建 ACL
	err = adminClient.CreateACLs([]*sarama.ResourceAcls{acl})
	if err != nil {
		log.Fatalf("Failed to create ACL: %v", err)
	} else {
		log.Println("ACL created successfully")
	}
}
