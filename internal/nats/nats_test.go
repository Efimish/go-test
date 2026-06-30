package nats

import (
	"fmt"
	"testing"
	"time"

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

func print(format string, a ...any) {
	fmt.Printf(
		"\033[32m[%s]\033[0m "+format,
		append(
			[]any{time.Now().Format("15:04:05.000000000")},
			a...,
		)...,
	)
}

func TestNatsSubscriberQueue(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("127.0.0.1", nats.Token("secret"))
	if err != nil {
		t.Errorf("Не удалось подключиться к серверу NATS: %s", err)
	}
	defer nc.Close()

	// Три подписчика
	subscribers := map[string]time.Duration{
		"1-🍎": 1000,
		"2-🍌": 1020,
		"3-🥝": 1200,
	}

	for name, ms := range subscribers {
		nc.QueueSubscribe("sub", "queue", func(m *nats.Msg) {
			print("%s) 🚀 Start: %s\n", name, m.Data)
			time.Sleep(time.Millisecond * ms)
			print("%s) 🎉 Finish: %s\n", name, m.Data)
		})
	}

	for i := 1; i <= 1000; i++ {
		// time.Sleep(time.Millisecond * 400)
		nc.Publish("sub", []byte(
			fmt.Sprintf("%d", i),
		))
	}
	select {}
}
