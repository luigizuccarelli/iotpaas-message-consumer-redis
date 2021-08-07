// +build fake

package connectors

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/microlib/simple"
)

var (
	// create a key value map (to fake redis)
	store map[string]string
)

type Connectors struct {
	Redis  *FakeRedis
	Logger *simple.Logger
	Kafka  *KafkaConsumerWrapper
	Flag   string
	Name   string
}

type FakeRedis struct {
}

type FakeKafkaConsumer struct {
	Flag string
	File string
}

type KafkaConsumerWrapper struct {
	Consumer *FakeKafkaConsumer
}

func (r *Connectors) Error(msg string, val ...interface{}) {
	r.Logger.Error(fmt.Sprintf(msg, val...))
}

func (r *Connectors) Info(msg string, val ...interface{}) {
	r.Logger.Info(fmt.Sprintf(msg, val...))
}

func (r *Connectors) Debug(msg string, val ...interface{}) {
	r.Logger.Debug(fmt.Sprintf(msg, val...))
}

func (r *Connectors) Trace(msg string, val ...interface{}) {
	r.Logger.Trace(fmt.Sprintf(msg, val...))
}

func (r *Connectors) Get(key string) (string, error) {
	return store[key], nil
}

func (r *Connectors) Set(key string, value string, expr time.Duration) (string, error) {
	if r.Flag == "errorDB" {
		return "", errors.New("forced redis db error")
	}
	store[key] = value
	return value, nil
}

func (r *Connectors) Close() {
	store = nil
}

func (c *Connectors) KafkaConsumer() *KafkaConsumerWrapper {
	return c.Kafka
}

func (f *FakeKafkaConsumer) Close() {
}

func (f *FakeKafkaConsumer) ReadMessage(int) (*kafka.Message, error) {
	if f.Flag == "error" {
		return &kafka.Message{}, errors.New("forced read message error")
	}

	b, _ := ioutil.ReadFile(f.File)
	return &kafka.Message{Value: b}, nil
}

// NewTestClientConnectors - inject our test connectors
func NewTestClientConnectors(filename string, err string, logger *simple.Logger) Clients {
	store = map[string]string{"": ""}
	cw := &KafkaConsumerWrapper{Consumer: &FakeKafkaConsumer{File: filename, Flag: err}}
	conn := &Connectors{Kafka: cw, Redis: &FakeRedis{}, Logger: logger, Name: "test", Flag: err}
	return conn
}
