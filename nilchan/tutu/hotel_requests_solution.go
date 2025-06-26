package tutu

import (
	"fmt"
	"sync"
	"time"
)

// Есть поток данных, в виде идентификаторов отелей, для каждого отеля нужно сделать поисковый запрос (запрос выполняется
// минимум 500ms) и отправить результаты в другой поток.

type SearchResult struct {
	HotelID int
}

func main() {
	dataCh := make(chan int)
	resultCh := make(chan SearchResult)
	var wg sync.WaitGroup

	go func() {
		for i := 0; i <= 10; i++ {
			dataCh <- i
		}
		defer close(dataCh)
	}()

	for ch := range dataCh {
		go func() {
			res := search(ch)
			resultCh <- res
		}()
	}

	// Количество воркеров, которые будут обрабатывать hotelID
	const workerCount = 3
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for hotelID := range dataCh {
				result := search(hotelID)
				resultCh <- result
			}
		}()
	}

	// Горутина для закрытия resultCh после завершения всех воркеров
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Чтение и вывод результатов
	for result := range resultCh {
		fmt.Println("Search result:", result)
	}
}

func search(hotelID int) SearchResult {
	time.Sleep(time.Millisecond * 500)
	return SearchResult{
		HotelID: hotelID,
	}
}
