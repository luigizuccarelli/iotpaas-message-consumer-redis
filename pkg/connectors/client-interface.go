package connectors

import "time"

// Clients interface - the NewClientConnectors will return this struct
type Clients interface {
	Error(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Trace(string, ...interface{})
	Get(string) (string, error)
	Set(string, string, time.Duration) (string, error)
	KafkaConsumer() *KafkaConsumerWrapper
	Close()
}
