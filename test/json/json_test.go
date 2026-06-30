package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Notification struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func TestJson() {
	notification := Notification{
		ID:      1,
		Message: "проверка",
	}

	// 1. Создаем канал для передачи готовой JSON-строки
	ch := make(chan string)

	// 2. Создаем потоковый пайп (трубу) для связки Encoder и Reader
	pr, pw := io.Pipe()

	// Горутина 1: Кодируем структуру в поток (PipeWriter)
	go func() {
		defer pw.Close() // Закрываем писатель, чтобы читатель понял, что данные закончились

		// Кодируем напрямую в пайп
		if err := json.NewEncoder(pw).Encode(notification); err != nil {
			fmt.Println("Ошибка кодирования:", err)
		}
	}()

	// Горутина 2: Читаем из потока (PipeReader), преобразуем в строку и отправляем в канал
	go func() {
		defer pr.Close()

		var buf bytes.Buffer
		// Копируем данные из пайпа в буфер памяти
		if _, err := io.Copy(&buf, pr); err != nil {
			fmt.Println("Ошибка чтения из пайпа:", err)
			return
		}

		// Отправляем готовую строку в канал
		ch <- buf.String()
	}()

	jsonString := <-ch
	fmt.Print("Получено из канала (JSON-строка):\n", jsonString)
}
