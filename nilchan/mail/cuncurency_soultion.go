package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	Name string
}

// main менять нельзя
func main() {
	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eeee"}}))
}

func fetchByName(ctx context.Context, userName string) (int, error) {
	// Тут происходит сетевой подход, который по userName возвращает userID
	time.Sleep(10 * time.Millisecond) // имитация сетевого подхода
	return rand.Int() & 100000, nil
}

// Все изменения должны производится в данной функции
func Do(ctx context.Context, users []User) (map[string]int, error) {
	collected := make(map[string]int, len(users))
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	errChan := make(chan error, len(users)) // для сбора ошибок

	for _, user := range users {
		user := user // важно: создаём копию для горутины

		wg.Add(1)
		go func() {
			defer wg.Done()

			id, err := fetchByName(ctx, user.Name)
			if err != nil {
				errChan <- err
				return
			}

			mu.Lock()
			collected[user.Name] = id
			mu.Unlock()
		}()
	}

	wg.Wait()
	close(errChan)
	/*
		почему без горутины
		потому что ты проверяешь len(errChan) > 0 сразу, не дождавшись завершения горутин,
		а errChan в это время может быть ещё пустой, даже если ошибки появятся позже.
	*/

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return collected, nil
	// TODO необходимо реализовать конкурентные запросы для каждого юзера и вернуть результат
}
