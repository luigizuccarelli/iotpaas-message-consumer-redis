// +build fake

package handlers

import (
	"fmt"
	"os"
	"testing"

	"github.com/luigizuccarelli/iotpaas-message-consumer-redis/pkg/connectors"
	"github.com/microlib/simple"
)

func TestAll(t *testing.T) {

	logger := &simple.Logger{Level: "trace"}

	t.Run("Message consumer : should pass", func(t *testing.T) {
		conn := connectors.NewTestClientConnectors("../../tests/payload.json", "normal", logger)
		os.Setenv("REDIS_HOST", "redis.redis-ha.svc.cluster.local")
		os.Setenv("KAFKA_BROKERS", "my-cluster-kafka-brokers.apache-kafka.svc.cluster.local:9092")
		os.Setenv("LOG_LEVEL", "trace")
		os.Setenv("SERVER_PORT", "")
		os.Setenv("URL", "http://127.0.0.1:7001/")
		os.Setenv("TOPIC", "test")
		os.Setenv("TESTING", "true")
		os.Setenv("CONNECTOR", "none")
		// call and test our consumer.go
		err := Init(conn)
		if err != nil {
			t.Fatal(fmt.Sprintf("Expected pass got error %v", err))
		}
	})

	t.Run("Message consumer : should fail forced (readmessage)", func(t *testing.T) {
		conn := connectors.NewTestClientConnectors("../../tests/payload.json", "error", logger)
		os.Setenv("REDIS_HOST", "redis.redis-ha.svc.cluster.local")
		os.Setenv("KAFKA_BROKERS", "my-cluster-kafka-brokers.apache-kafka.svc.cluster.local:9092")
		os.Setenv("LOG_LEVEL", "trace")
		os.Setenv("SERVER_PORT", "")
		os.Setenv("URL", "http://127.0.0.1:7001/")
		os.Setenv("TOPIC", "test")
		os.Setenv("TESTING", "true")
		os.Setenv("CONNECTOR", "none")
		// call and test our consumer.go
		err := Init(conn)
		if err == nil {
			t.Fatal("Expected fail got nil")
		}
	})

	t.Run("Message consumer : should fail (json parse)", func(t *testing.T) {
		conn := connectors.NewTestClientConnectors("../../tests/error.json", "normal", logger)
		os.Setenv("REDIS_HOST", "redis.redis-ha.svc.cluster.local")
		os.Setenv("KAFKA_BROKERS", "my-cluster-kafka-brokers.apache-kafka.svc.cluster.local:9092")
		os.Setenv("LOG_LEVEL", "trace")
		os.Setenv("SERVER_PORT", "")
		os.Setenv("URL", "http://127.0.0.1:7001/")
		os.Setenv("TOPIC", "test")
		os.Setenv("TESTING", "true")
		os.Setenv("CONNECTOR", "none")
		// call and test our consumer.go
		err := Init(conn)
		if err == nil {
			t.Fatal("Expected fail got nil")
		}
	})

	t.Run("Message consumer : should fail (forced db error)", func(t *testing.T) {
		conn := connectors.NewTestClientConnectors("../../tests/payload.json", "errorDB", logger)
		os.Setenv("REDIS_HOST", "redis.myportfolio.svc.cluster.local")
		os.Setenv("KAFKA_BROKERS", "my-cluster-kafka-brokers.apache-kafka.svc.cluster.local:9092")
		os.Setenv("LOG_LEVEL", "trace")
		os.Setenv("SERVER_PORT", "")
		os.Setenv("URL", "http://127.0.0.1:7001/")
		os.Setenv("TOPIC", "test")
		os.Setenv("TESTING", "true")
		os.Setenv("CONNECTOR", "none")
		// call and test our consumer.go
		err := Init(conn)
		if err == nil {
			t.Fatal("Expected fail got nil")
		}
	})
}
