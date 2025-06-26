package concurency

import (
	"context"
	"sync"
	"time"
)

//===========================================================
//Задача 7
//===========================================================

type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

/*
Пояснение:
Все функции запускаются параллельно.
Как только одна из них успешно возвращает результат — он отдается вызывающему коду.
Ошибки накапливаются, последняя возвращается, если все SearchFunc завершились с ошибками.
ctx можно использовать для отмены всех запросов (если передан извне, например с таймаутом).
Дополнительно можно убрать time.After(...), если не нужно ограничивать по времени помимо ctx.
*/

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
	// Когда получаем первый успешный результат - отдаем его сразу. Если все SearchFunc отработали
	// с ошибкой - отдаем последнюю полученную ошибку

	errCh := make(chan error, len(sfs))
	results := make(chan Result, len(sfs))

	wg := sync.WaitGroup{}
	wg.Add(len(sfs))

	for _, sf := range sfs {
		go func(sf SearchFunc) {
			defer wg.Done()

			res, err := sf(ctx, query)
			if err != nil {
				errCh <- err
				return
			}

			// отправляем результат, но только если канал открыт
			select {
			case results <- res:
			default:
			}
		}(sf)
	}

	// Закрываем каналы, когда все горутины завершатся
	go func() {
		wg.Wait()
		defer close(errCh)
		defer close(results)
	}()

	// ждем первый результат или все ошибки
	select {
	case res := <-results:
		return res, nil
	case <-ctx.Done():
		return Result{}, ctx.Err()
	case <-time.After(100 * time.Millisecond): // safety net timeout, опционально
	}

	// Все вернули ошибку — возвращаем последнюю
	var lastErr error
	for err := range errCh {
		lastErr = err
	}

	return Result{}, lastErr
}
