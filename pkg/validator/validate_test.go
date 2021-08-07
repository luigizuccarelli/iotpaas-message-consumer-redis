package validator

import (
	"fmt"
	"os"
	"testing"

	"github.com/microlib/simple"
)

func TestEnvars(t *testing.T) {
	logger := &simple.Logger{Level: "info"}

	t.Run("ValidateEnvars : should fail", func(t *testing.T) {
		os.Setenv("SERVER_PORT", "")
		err := ValidateEnvars(logger)
		if err == nil {
			t.Errorf(fmt.Sprintf("Handler %s returned with no error - got (%v) wanted (%v)", "ValidateEnvars", err, nil))
		}
	})

	t.Run("ValidateEnvars : should pass", func(t *testing.T) {
		os.Setenv("LOG_LEVEL", "info")
		os.Setenv("SERVER_PORT", "9000")
		os.Setenv("REDIS_HOST", "127.0.0.1:6379")
		os.Setenv("REDIS_PASSWORD", "test")
		os.Setenv("TOKEN", "dsafsdfdsf")
		os.Setenv("VERSION", "1.0.3")
		os.Setenv("KAFKA_BROKERS", "localhost:9092")
		os.Setenv("TOPIC", "test")
		os.Setenv("CONNECTOR", "NA")
		os.Setenv("NAME", "test")
		err := ValidateEnvars(logger)
		if err != nil {
			t.Errorf(fmt.Sprintf("Handler %s returned with error - got (%v) wanted (%v)", "ValidateEnvars", err, nil))
		}
	})

}
