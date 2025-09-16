package ЦУМ

import (
	"context"
	"sync"
	"time"
)

type Result struct {
	msg string
	err error
}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
	// Когда получает первый успешный результат - отдает его сразу, не дожидаясь результата других SearchFunc
	// Если все SearchFunc отработали с ошибкой - отдает последнюю полученную ошибку

	results := make(chan Result, len(sfs))
	errCh := make(chan error, len(sfs))

	wg := sync.WaitGroup{}
	wg.Add(len(sfs))

	for _, sf := range sfs {
		go func(sf SearchFunc) {
			defer wg.Done()

			resp, err := sf(ctx, query)
			if err != nil {
				errCh <- err
				return
			}

			// отправляем результат, но только если канал открыт
			select {
			case results <- resp: // отправим, если можем
			default: // иначе просто игнорируем
			}

		}(sf)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errCh)
	}()

	// ждем первый результат или все ошибки
	select {
	case res := <-results:
		return res, nil
	case <-ctx.Done():
		return Result{}, ctx.Err()
	case <-time.After(100 * time.Millisecond):
	}

	// Все вернули ошибку — возвращаем последнюю
	var lastErr error
	for err := range errCh {
		lastErr = err
	}

	return Result{}, lastErr
}
