package concurency

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//===========================================================
//Задача 14
//1. Что выведет на экран и сколько времени будет работать?
//2. Нужно ускорить, чтобы работало быстрее. Сколько будет работать теперь?
//3. Если бы в networkRequest выполнялся реальный сетевой вызов, то какие с какими проблемами мы могли бы столкнуться в данном коде?
//4. Если url немного, а запросов к ним много, то как можно оптимизировать?
//===========================================================

/*
1. Каждый запрос занимает ~1 мс (time.Sleep(time.Millisecond))
Запросы выполняются параллельно (горутины)
Основное ограничение - количество CPU и contention на мьютексе
Ориентировочное время: 100-500 мс (зависит от CPU)
*/

/*
2. Как ускорить?
Проблемы текущей реализации:
Мьютекс создает contention при инкременте счетчика
Слишком много горутин (10k) создают overhead
Улучшения:
atomic.Int32 вместо мьютекса
Семафор для ограничения параллелизма (100 горутин)
Уменьшение накладных расходов на создание горутин
*/

/*
3. Проблемы с реальными сетевыми запросами
Основные проблемы:
Нет таймаутов:
Зависший запрос заблокирует горутину навсегда
Решение: context.WithTimeout
Нет обработки ошибок:
Ошибки сети просто игнорируются
Решение: возвращать ошибки из networkRequest
Нет retry логики:
Временные сбои приведут к потере запросов
Решение: добавить retry с экспоненциальным backoff
Нет ограничения скорости:
Можно перегрузить сервер
Решение: rate limiting (например, golang.org/x/time/rate)
Нет circuit breaker:
При проблемах с сервером продолжит посылать запросы
Решение: добавить паттерн Circuit Breaker
*/

/*
Если URL повторяются:
Кеширование ответов:
var cache = sync.Map{}
func networkRequest(url string) {
	if val, ok := cache.Load(url); ok {
		return val
	}
	// ... выполняем запрос
	cache.Store(url, result)
}
Пул соединений:
Использовать http.Client с настроенным Transport
Поддерживать keep-alive соединения
Пакетные запросы:
Объединять одинаковые запросы в batch
Пример: вместо 100 запросов к /user/1, сделать 1 запрос к /users?ids=1,2,3...
Шардирование запросов:
var urlShards = make([]chan string, 10)
// Каждая горутина обрабатывает свой shard
*/

const numRequests = 10000

var count atomic.Int32 // Атомарный счетчик
// var count int
var cache sync.Map
var m sync.Mutex

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	m.Lock()
	count++
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(numRequests)

	sem := make(chan struct{}, 100) // Ограничиваем параллелизм
	urls := generateUrls()          // []string с повторениями

	for _, url := range urls {
		sem <- struct{}{}
		go func(u string) {
			defer func() { <-sem }()

			if _, ok := cache.Load(u); ok {
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			if err := realNetworkRequest(ctx, u); err == nil {
				cache.Store(u, true)
				count.Add(1)
			}
		}(url)
	}

	// Вариант 1
	for i := 0; i < numRequests; i++ {
		sem <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-sem
			}()
			networkRequest()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
