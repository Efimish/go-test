package nats

import (
	"fmt"
	"sync"
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

	// Simple Async Subscriber 1
	nc.Subscribe("subject1", func(m *nats.Msg) {
		t.Logf("Received a message (1): %s\n", string(m.Data))
	})

	// Simple Async Subscriber 2
	nc.Subscribe("subject1", func(m *nats.Msg) {
		t.Logf("Received a message (2): %s\n", string(m.Data))
	})

	// Simple Publishers
	nc.Publish("subject1", []byte("1: Hello World\r\nHello\r\n\r\n\r\nTest"))
	nc.Publish("subject1", []byte("2: Hello World"))
	nc.Publish("subject1", []byte("3: Hello World"))
	nc.Publish("subject1", []byte("4: Hello World"))
	nc.Publish("subject1", []byte("5: Hello World"))

	select {}
}

func TestBeforeTest(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Errorf("Не удалось подключиться к серверу NATS: %s", err)
	}
	defer nc.Close()

	// nc.Subscribe("sub", func(m *nats.Msg) {
	// 	t.Logf("Received a message: %s\n", string(m.Data))
	// })
	nc.QueueSubscribe("sub", "queue", func(m *nats.Msg) {
		t.Logf("Sub 1) Received message: %s", m.Data)
	})

	for i := 1; i <= 10; i++ {
		nc.Publish("sub", []byte(
			fmt.Sprintf("%d) Message!", i),
		))
	}

	select {}
}

func TestNatsSubscriberQueue(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Errorf("Не удалось подключиться к серверу NATS: %s", err)
	}
	defer nc.Close()

	// Use a WaitGroup to wait for 10 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(10)

	// Create a queue subscription on "updates" with queue name "workers"
	if _, err := nc.QueueSubscribe("updates", "workers", func(m *nats.Msg) {
		wg.Done()
	}); err != nil {
		t.Fatal(err)
	}
}
