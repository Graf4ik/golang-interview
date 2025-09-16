package МТС

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Semaphore
// У вас есть список URL-адресов, которые нужно обработать параллельно.
// Однако одновременно можно запускать не более N горутин. Вместо реального
// выполнения HTTP-запросов использовать httpGet()
// Реализуйте функцию fetchURL , которая:
// 1. Захватывает семафор перед началом работы.
// 2. Добавляет результат в общий срез results .
// 3. Освобождает семафор после завершения.

func httpGet(url string) string {
	time.Sleep(time.Second)
	return fmt.Sprintf("%s: 200 OK", url)
}

func fetchURL(url string, semaphore chan struct{}, results *[]string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	semaphore <- struct{}{}

	// Захватываем семафор.
	defer func() {
		<-semaphore // Освобождаем семафор в конце работы.
	}()

	// Эмуляция долгого запроса.
	time.Sleep(2 * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Добавляем результат в общий массив
	mu.Lock()
	*results = append(*results, fmt.Sprintf("URL: %s, Status: %s", url, resp.Status))
	mu.Unlock()
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://golang.org",
	}
	const maxConcurrency = 2
	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, maxConcurrency)
	results := make([]string, len(urls))
	mu := sync.Mutex{}

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, semaphore, &results, &wg, &mu)
	}

	wg.Wait()
	fmt.Println("Results:")
	for _, result := range results {
		fmt.Println(result)
	}
}
