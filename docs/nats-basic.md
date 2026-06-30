# Basic NATS example

```go
package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Подключение к серверу NATS
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу NATS: %s", err)
	}
	defer nc.Close()

	// Подписчик
	nc.Subscribe("subject1", func(m *nats.Msg) {
		fmt.Printf("Получено сообщение: %s\n", string(m.Data))
	})

	// Отправитель
	nc.Publish("subject1", []byte("Привет через NATS"))

	// Ожидание других потоков
	select {}
}
```
