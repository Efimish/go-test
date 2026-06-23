package nats

import (
	"testing"

	"github.com/nats-io/nats.go"
)

func TestNats(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Errorf("Не удалось подключиться к серверу NATS: %s", err)
	}
	defer nc.Close()

	// Simple Async Subscriber
	nc.Subscribe("subject1", func(m *nats.Msg) {
		t.Logf("Received a message: %s\n", string(m.Data))
	})

	// Simple Publisher
	nc.Publish("subject1", []byte("Hello World"))
}
