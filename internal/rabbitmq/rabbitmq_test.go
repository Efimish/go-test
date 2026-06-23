package rabbitmq

import (
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

var url = "amqp://127.0.0.1:5672"

func TestRabbitMQ(t *testing.T) {
	// Подключение к RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		t.Errorf("Не удалось подключиться к RabbitMQ: %s\n", err)
	}
	defer conn.Close()

	// Создание канала RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		t.Errorf("Не удалось открыть канал в RabbitMQ: %s\n", err)
	}
	defer ch.Close()

	// Создание очереди RabbitMQ
	q, err := ch.QueueDeclare(
		"notifications",
		true,
		false,
		false,
		false,
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	if err != nil {
		t.Errorf("Не удалось создать очередь в RabbitMQ: %s\n", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		t.Errorf("Не удалось зарегистрировать получателя: %s\n", err)
	}

	t.Logf(" [*] Ожидание сообщений. Для выхода нажмите CTRL+C")
	go func() {
		for d := range msgs {
			t.Logf("Получено сообщение: %s", d.Body)
		}
	}()

	ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("проверка"),
	})

	t.Logf("Очередь создана: %+v\n", q)
}
