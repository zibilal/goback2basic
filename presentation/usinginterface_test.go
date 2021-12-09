package presentation

import (
	"fmt"
	"testing"
)

type Logger interface {
	Log(json string) error
}

type Producer interface {
	Produce(dt string) error
}

type KafkaProducer struct {}
func (k KafkaProducer) Produce(dt string) error {
	return nil
}

type KafkaLogger struct {
	sender Producer
}

func(k KafkaLogger) Log(j string) error {
	return k.sender.Produce(j)
}

func ALog(msg string, l Logger) error {

	if err := l.Log(msg); err != nil {
		fmt.Printf("error is happening, log not displayed, due to %v", err)
		return err
	}
	return nil
}

func TestUsingInterface(t *testing.T) {
	logger := KafkaLogger{
		sender: KafkaProducer{},
	}
	err := ALog("Process A started", logger)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	} else {
		t.Log("Succes")
	}
}


