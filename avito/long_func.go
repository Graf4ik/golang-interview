package avito

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.

func LongFunc() int {
	time.Sleep(2 * time.Second) //  имитация долгой работы
	return 42
}

func LongFuncWithTimeout(timeout time.Duration) (int, error) {
	start := time.Now()
	resultChan := make(chan int, 1)

	go func() {
		result := LongFunc()
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		duration := time.Since(start)
		log.Printf("Function completed in %v", duration)
		return res, nil
	case <-time.After(timeout):
		duration := time.Since(start)
		log.Printf("timeout after %s", duration)
		return 0, fmt.Errorf("function timeout after %s", timeout)
	}
}

// Реализация через context.WithTimeout
func LongFuncWithContextTimeout2(timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	start := time.Now()
	resultChan := make(chan int, 1)

	go func() {
		res := LongFunc() // выполняется независимо
		select {
		case resultChan <- res: // пытаемся отправить результат
		case <-ctx.Done(): // но если контекст отменился — просто выходим
		}
	}()

	select {
	case res := <-resultChan:
		log.Printf("Function completed in %v", time.Since(start))
		return res, nil
	case <-ctx.Done():
		log.Printf("Function timed out after %v", time.Since(start))
		return 0, ctx.Err()
	}
}

/*
🔍 Объяснение:
Используем time.After(timeout) — создаёт канал, который «стреляет» по истечении таймаута.
Запускаем функцию в отдельной горутине и читаем её результат через канал.
Логируем, сколько прошло времени в обоих случаях (успех и таймаут).
Если функция "зависла" — возвращаем ошибку, игнорируя результат.
*/

func main() {
	res, err := LongFuncWithTimeout(1 * time.Second)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Result: %d", res)
	}
}
