package main

import (
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
)

//===========================================================
//Задача 13
//1. Запросить параллельно данные из источников. Если все где-то произошла ошибка, то вернуть ошибку, иначе вернуть nil.
//2. Представим, что теперь функция должна возвращать результат int. Есть функция resp.Size(), для каждого url
//надо проссумировать и вернуть, если ошибок не было. Просто описать подход к решению
//3. Что делать, если урлов у нас миллионы?
//
//===========================================================

func main() {
	_, err := download1([]string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
	})

	if err != nil {
		panic(err)
	}
}

const maxWorkers = 100

/*
Параллельно запросить данные из источников. Если все запросы вернули ошибку — вернуть ошибку, иначе nil.
Подход:
Запустить параллельные запросы (через горутины).
Использовать sync.WaitGroup для ожидания завершения всех.
Считать количество ошибок, желательно через atomic.
Если все запросы завершились с ошибкой — вернуть ошибку.
*/

/*
Теперь download() должна возвращать int — сумму resp.Size(), если не было ошибок.
Подход:
Предполагаем, что resp.Size() — это размер тела ответа (например, len(body)).
Создаём totalSize int32/int64, безопасно увеличиваем его (atomic.AddInt64()).
Если какая-либо ошибка — возвращаем её.
Если все с ошибкой — вернуть ошибку, иначе вернуть сумму.
*/

func download1(urls []string) (int64, error) {
	var wg sync.WaitGroup
	var errorCount int32
	var totalSize int64

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != http.StatusOK {
				atomic.AddInt32(&errorCount, 1)
				return
			}
			defer resp.Body.Close()

			size := resp.ContentLength
			if size > 0 {
				atomic.AddInt64(&totalSize, size)
			}
		}(url)
	}

	wg.Wait()

	if int(errorCount) == len(urls) {
		return 0, errors.New("all requests failed")
	}

	return totalSize, nil
}

/*
Что делать, если URL-ов миллионы?
Подход:
Ограничение параллелизма:
Создать семафор (например, workerPool := make(chan struct{}, maxWorkers)), чтобы не запускать миллионы горутин.
Обработка по чанкам или через очередь:
Использовать worker pool + канал задач, чтобы эффективно и стабильно обрабатывать огромный список URL-ов.
Возможен подход с batching + временными файлами, если нужно сохранять данные (например, response body) и не хватает оперативной памяти.
*/

func download2(urls []string) (int64, error) {
	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	var totalSize int64
	var errorCount int32

	jobs := make(chan string, len(urls))

	// Воркеры
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				resp, err := http.Get(url)
				if err != nil || resp.StatusCode != http.StatusOK {
					atomic.AddInt32(&errorCount, 1)
					continue
				}
				size := resp.ContentLength
				if size > 0 {
					atomic.AddInt64(&totalSize, size)
				}
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	if int(errorCount) == len(urls) {
		return 0, errors.New("all requests failed")
	}

	return totalSize, nil
}
