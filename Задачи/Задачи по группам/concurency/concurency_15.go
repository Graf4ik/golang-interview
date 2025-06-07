package main

import (
	"fmt"
	"math/rand"
	"time"
)

//===========================================================
//Задача 15
//
//===========================================================

// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).

// Нужно написать обертку predictableFunc,
// которая будет работать с заданным фиксированным таймаутом (например, 1 секунду).

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.

// Вариант 1
func predictableFunc() (int64, error) {
	resultChan := make(chan int64, 1)
	start := time.Now()
	go func() {
		resultChan <- unpredictableFunc()
	}()

	select {
	case result := <-resultChan:
		elapsed := time.Since(start)
		fmt.Printf("Function executed in %v\n", elapsed)
		return result, nil
	case <-time.After(1 * time.Second):
		return 0, fmt.Errorf("timeout after 1 second")
	}
}

// Вариант 2
//func predictableFunc() (int64, error) {
//	// Создаем контекст с таймаутом 1 секунда
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel() // Освобождаем ресурсы контекста
//
//	// Канал для результата
//	resultChan := make(chan int64, 1)
//
//	// Запускаем функцию в горутине, передавая контекст
//	start := time.Now()
//	go func() {
//		resultChan <- unpredictableFunc()
//	}()
//
//	// Ожидаем результат или отмену контекста
//	select {
//	case result := <-resultChan:
//		elapsed := time.Since(start)
//		fmt.Printf("Function executed in %v\n", elapsed)
//		return result, nil
//	case <-ctx.Done():
//		return 0, ctx.Err() // Возвращаем ошибку таймаута
//	}
//}

func main() {
	fmt.Println("started")

	if result, err := predictableFunc(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
