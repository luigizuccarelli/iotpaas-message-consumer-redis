// +build real

package connectors

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis"
	"github.com/microlib/simple"
)

// The premise here is to use this as a reciever in the relevant functions
// this allows us then to mock/fake connections and calls
type Connectors struct {
	Logger *simple.Logger
	Kafka  *KafkaConsumerWrapper
	Redis  *redis.Client
	Name   string
}

type KafkaConsumerWrapper struct {
	Consumer *kafka.Consumer
}

// NewClientConnectors : function that initialises connections to DB's, caches' queues etc
// Seperating this functionality here allows us to inject a fake or mock connection object for testing
func NewClientConnectors(logger *simple.Logger) Clients {

	logger.Info(fmt.Sprintf("Redis envars: %s %s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD")))

	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_HOST"),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		//Password:     os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	logger.Info(fmt.Sprintf("Redis connection: %v", redisClient))

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKERS"),
		"group.id":          "iotpaasGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{os.Getenv("TOPIC")}, nil)
	logger.Info(fmt.Sprintf("Kafka brokers: %s", os.Getenv("KAFKA_BROKERS")))

	cw := &KafkaConsumerWrapper{Consumer: c}
	return &Connectors{Redis: redisClient, Kafka: cw, Logger: logger, Name: "RealConnectors"}
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

// Get - redis function wrapper
func (r *Connectors) Get(key string) (string, error) {
	val, err := r.Redis.Get(key).Result()
	return val, err
}

// Set - redis function wrapper
func (r *Connectors) Set(key string, value string, expr time.Duration) (string, error) {
	val, err := r.Redis.Set(key, value, expr).Result()
	return val, err
}

// KafkConsumer - function wrapper
func (c *Connectors) KafkaConsumer() *KafkaConsumerWrapper {
	return c.Kafka
}

func (c *Connectors) Close() {
	c.Redis.Close()
	c.Kafka.Consumer.Close()
}
