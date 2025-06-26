package fabricSolutions

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// worker pool:
// организовать запуск не более 3 параллельных воркеров, которые выполняют запросы
// на все url из массива urls
// пример урлов, которые получают из генератора:

// http://echo.xff.pw:8080/?echo_time=100&echo_body=body-1
// http://echo.xff.pw:8080/?echo_time=100&echo_body=body-2
// http://echo.xff.pw:8080/?echo_time=100&echo_body=body-3
// http://echo.xff.pw:8080/?echo_time=100&echo_body=body-4
// http://echo.xff.pw:8080/?echo_time=100&echo_body=body-5

// Функцию желательно никак не менять, но если требуется по реализации - можно

func main() {
	pool := makePool(3)
	for url := range generateUrls(5) {
		log.Printf("=> Запуск, %s", url)
		pool.DoWork(url)
	}
	pool.Wait()
}

type Pool struct {
	tasks chan string
	wg    sync.WaitGroup
}

func makePool(size int) *Pool {
	p := &Pool{
		tasks: make(chan string, size),
	}

	for i := 0; i < size; i++ {
		go func(id string) {
			for url := range p.tasks {
				resp, err := http.Get(url)
				if err != nil {
					log.Println(err)
				} else {
					body, _ := io.ReadAll(resp.Body)
					resp.Body.Close()
					log.Printf("[worker-%d] Ответ: %s", id, string(body))
				}
				p.wg.Done()
			}
		}(strconv.Itoa(i))
	}

	return p
}

func (p *Pool) DoWork(url string) {
	p.wg.Add(1)
	p.tasks <- url
}

func (p *Pool) Wait() {
	p.wg.Wait()
	close(p.tasks)
}

func generateUrls(size int) chan string {
	out := make(chan string, size)
	go func() {
		for i := 0; i < size; i++ {
			out <- "http://echo.xff.pw:8080/?echo_time=100&echo_body=body-" + string('0'+i)
		}
	}()
	return out
}
