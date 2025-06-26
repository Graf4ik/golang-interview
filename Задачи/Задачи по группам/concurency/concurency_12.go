package concurency

import (
	"fmt"
	"sync"
)

//===========================================================
//Задача 12
//1. Конурентно по батчам запросить данные и записать в файл. Нужна общая конструкция, функции которые делают запрос к сайту и выгрузку в файл можно не реализовывать.
//2. Сделать так, чтобы одновременно выполнялось не более chunkSize запросов.
//===========================================================
/*
 Пояснение:
sem := make(chan struct{}, chunkSize) — ограничивает количество одновременных горутин (и, соответственно, запросов).
worker — запускается для каждого id, делает запрос и сохраняет результат.
wg — синхронизация завершения всех воркеров.
fetchData и saveData — заглушки, которые ты можешь заменить на реальную логику.

*/
// Константы
const (
	urlTemplate = "http://jsonplaceholder.typicode.com/tools/%d"
	chunkSize   = 100
	dataCount   = 2 << 10 // 2048
)

// Заглушка для запроса (эмуляция)
func fetchData(id int) (string, error) {
	// Здесь можно сделать реальный HTTP-запрос, например:
	// resp, err := http.Get(fmt.Sprintf(urlTemplate, id))
	// defer resp.Body.Close()
	// Но пока просто возвращаем заглушку:
	return fmt.Sprintf("Data for ID %d", id), nil
}

// Заглушка для записи в файл (можно сделать буфер и потом слить в файл)
func saveData(id int, data string) error {
	// Здесь можно реализовать запись в файл или канал
	fmt.Printf("Saving ID %d: %s\n", id, data)
	return nil
}

func worker(id int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	// Ограничение количества одновременно работающих горутин
	sem <- struct{}{}        // блокируем
	defer func() { <-sem }() // разблокируем

	// Запрос данных
	data, err := fetchData(id)
	if err != nil {
		fmt.Printf("Error fetching ID %d: %v\n", id, err)
		return
	}

	// Сохранение
	err = saveData(id, data)
	if err != nil {
		fmt.Printf("Error saving ID %d: %v\n", id, err)
	}
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, chunkSize) // семафор на chunkSize одновременных запросов

	for i := 1; i <= dataCount; i++ {
		wg.Add(1)
		go worker(i, &wg, sem)
	}

	wg.Wait()
	fmt.Println("All done.")
}
