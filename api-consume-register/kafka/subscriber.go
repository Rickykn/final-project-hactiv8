package kafka

import (
	"api-consume-register/host"
	"api-consume-register/models"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

const (
	zookeeperConn = "localhost:2181"
	consGroup     = "group1"
	topic         = "register"
)

func ConnConsumer() (*consumergroup.ConsumerGroup, error) {
	config := consumergroup.NewConfig()

	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	consGroup, err := consumergroup.JoinConsumerGroup(consGroup, []string{topic}, []string{zookeeperConn}, config)
	if err != nil {
		os.Exit(1)
	}
	return consGroup, nil
}

func Subscriber(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			if msg.Topic != topic {
				continue
			}

			fmt.Println("Topic : ", msg.Topic)
			fmt.Println("Topic : ", string(msg.Value))
			go sendCreateAccount(msg.Value)
			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper :", err.Error())
			}
		}
	}
}

func sendCreateAccount(msg []byte) {
	var register models.Person
	err := json.Unmarshal(msg, &register)
	if err != nil {
		fmt.Println("Message Broken")
	}

	acc := host.Account{
		Name:         register.Name,
		Email:        register.Email,
		Number_phone: register.Number_phone,
	}

	user := models.Person{
		Email:        register.Email,
		Name:         register.Name,
		Address:      register.Address,
		Number_phone: register.Number_phone,
	}

	go host.SendCreateAccount(acc)
	go host.SendCreateUser(user)
}
