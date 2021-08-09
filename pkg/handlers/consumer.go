package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/luigizuccarelli/iotpaas-message-consumer-redis/pkg/connectors"
	"github.com/luigizuccarelli/iotpaas-message-consumer-redis/pkg/schema"
)

// Init : public function that connects to the kafka queue and redis cache
func Init(conn connectors.Clients) error {

	var run bool = true
	var err error
	var msg *kafka.Message
	cw := conn.KafkaConsumer()

	for run == true {
		msg, err = cw.Consumer.ReadMessage(-1)
		if err == nil {
			err = postToDB(conn, msg)
		} else {
			// The client will automatically try to recover from all errors.
			conn.Error("Consumer error: %v (%v)\n", err, msg)
		}
		test, _ := strconv.ParseBool(os.Getenv("TESTING"))
		if test == true {
			run = false
		}
	}

	cw.Consumer.Close()
	return err
}

// postToDB : private utility function that posts the json payload to couchbase
func postToDB(conn connectors.Clients, msg *kafka.Message) error {

	var data *schema.IOTPaaS

	if msg != nil {

		errs := json.Unmarshal(msg.Value, &data)
		if errs != nil {
			conn.Error("postToDB unmarshalling format %v", errs)
			return errs
		}

		_, err := conn.Set(data.Id, string(msg.Value), -1)
		if err != nil {
			conn.Error(fmt.Sprintf("Could not insert schema into redis %v", err))
			return err
		}

		// all good :)
		return nil

	} else {
		conn.Info("Message data is nil")
		return nil
	}
}
