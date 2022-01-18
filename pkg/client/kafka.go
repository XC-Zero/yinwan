package client

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

const MaxMessageBytes = 4 * 1024 * 1024

var kafkaInstance sarama.SyncProducer

func InitKafka(config cfg.KafkaConfig) {
	cfg := sarama.NewConfig()
	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}
	cfg.ClientID = hostname
	cfg.Producer.MaxMessageBytes = MaxMessageBytes
	cfg.Producer.Return.Successes = true
	if config.Username == "" || config.Password == "" {
		panic(fmt.Sprintf("KafkaConfig not have username or password! Error is  %s", err.Error()))
	}
	cfg.Net.SASL.Enable = true
	cfg.Net.SASL.User = config.Username
	cfg.Net.SASL.Password = config.Password
	address := config.AddrList

	pro, err := sarama.NewSyncProducer(address, cfg)
	if err != nil {
		panic(fmt.Sprintf("初始化kafka失败: %s", err.Error()))
	}
	client, err := sarama.NewClient(address, cfg)
	if err != nil {
		panic(fmt.Sprintf("初始化kafka失败: %s", err.Error()))
	}
	KafkaClient = &client

	kafkaInstance = pro
}

// PushInterfaceToKafka 无限定格式推送kafka
func PushInterfaceToKafka(topic string, data []interface{}) error {
	var messages []*sarama.ProducerMessage
	for i := range data {
		// 为了方便阅读，加个 \t
		bytes, err := json.MarshalIndent(data[i], "", "\t")
		if err != nil {
			log.Println(err)
		}
		messages = append(messages, &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(bytes),
		})
	}
	errs := kafkaInstance.SendMessages(messages)
	if errs != nil {
		pde, ok := errs.(sarama.ProducerErrors)
		if !ok {
			return errors.New("send msg to kafka fail:" + errs.Error())
		}
		errorMsg := ""
		for i := 0; i < len(pde); i++ {
			if strings.Contains(pde[i].Error(), "circuit breaker is open") {
				continue
			}
			errorMsg += "send msg to kafka fail:" + pde[i].Error()
		}
		return errors.New(errorMsg)

	}
	return nil
}

func CloseKafka() {
	err := kafkaInstance.Close()
	if err != nil {
		log.Println(err)
	}
}
