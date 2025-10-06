package messagebroker

import (
	"context"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

var instanceKafkaReader *kafka.Reader
var instanceKafkaWrite *kafka.Writer

var once sync.Once

func GetKafkaReaderMessageBroken() *kafka.Reader {
	once.Do(func() {
		// Configuração do reader (consumer)
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"kafka:9092"},
			Topic:   "meu-topico",
			GroupID: "meu-grupo",
		})
		instanceKafkaReader = reader
	})
	return instanceKafkaReader
}

func GetKafkaWriterMessageBroken() *kafka.Writer {
	once.Do(func() {
		// Configuração do reader (consumer)
		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{"kafka:9092"},
			Topic:   "meu-topico",
		})
		instanceKafkaWrite = writer
	})
	return instanceKafkaWrite
}

func Producer(key string, message string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
		Time:  time.Now(),
	}

	return GetKafkaWriterMessageBroken().WriteMessages(context.Background(), msg)
}

func Consumer() (string, error) {
	m, err := GetKafkaReaderMessageBroken().ReadMessage(context.Background())
	if err != nil {
		return "", err
	}
	return string(m.Value), nil
}
