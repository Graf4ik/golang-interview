package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	err := predictableTimeWork()
	if err != nil {
		return
	}
}

// написать обертку для этой функции, которая будет прерывать выполнение,
// если функция работает больше 3 секунд, и возвращать ошибку
func predictableTimeWork() error {
	ch := make(chan struct{})

	go func() {
		randomWork()
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-time.After(3 * time.Second):
		return fmt.Errorf("error")
	}
}

// написать функцию, которая работает неопределенно долго (до 100 секунд)
func randomWork() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
