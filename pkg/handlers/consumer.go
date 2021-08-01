package handlers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	gocb "github.com/couchbase/gocb/v2"
	"github.com/luigizuccarelli/iotpaas-message-consumer/pkg/connectors"
	"github.com/luigizuccarelli/iotpaas-message-consumer/pkg/schema"
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

	// check if we have the updated detached json
	if msg != nil {
		payload, _ := url.PathUnescape(string(msg.Value))
		conn.Trace(fmt.Sprintf("Data from message queue %s", payload))

		// we have the new format
		errs := json.Unmarshal(msg.Value, &data)
		if errs != nil {
			conn.Error("postToDB unmarshalling new format %v", errs)
			return errs
		}

		conn.Debug(fmt.Sprintf("IOTPaaS struct  %v", data))
		res, err := conn.Upsert(data.DeviceId, data, &gocb.UpsertOptions{})
		conn.Debug(fmt.Sprintf("IOTPaaS from insert %v", res))
		if err != nil {
			conn.Error(fmt.Sprintf("Could not insert schema into couchbase %v", err))
			return err
		}

		// all good :)
		conn.Info("IOTPaas schema inserted into couchbase")
		return nil

	} else {
		conn.Info("Message data is nil")
		return nil
	}
}
