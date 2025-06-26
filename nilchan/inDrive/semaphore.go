package inDrive

import (
	"fmt"
	"sync"
	"time"
)

// Необходимо дописать код, чтобы задача выполнялась асинхронно с ограниченем по количеству одновременных запросов

const maxConcurrent = 3

func main() {
	semaphore := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for _, val := range getData() {
		semaphore <- struct{}{} // занять слот (блокирует, если лимит исчерпан)
		wg.Add(1)

		go func(h string) {
			defer func() {
				<-semaphore // освободить слот
				wg.Done()   // сообщить WaitGroup о завершении
			}()

			if !checkDomain(val) {
				fmt.Printf("%s is bad \n", val)
			} else {
				fmt.Printf("%s is good \n", val)
			}
		}(val)
	}

	go func() {
		close(semaphore)
		wg.Wait()
	}()
}

func getData() []string {
	return []string{"1.test", "2.test", "3.test", "4.test", "5.test", "6.test", "7.test"}
}

func checkDomain(host string) bool {
	// имитация сетевого запроса
	time.Sleep(500 * time.Millisecond)
	return host[0]%3 != 0
}

/*
Как работает:

sem — буферизированный канал-семафор размером maxConcurrent.
Отправка в него («sem <- struct{}{}») блокируется, если в канале уже maxConcurrent элементов.
Каждая горутина после работы делает <-sem, освобождая «слот».
WaitGroup гарантирует, что main завершится лишь после всех проверок.
Порядок вывода может быть любым, но одновременно выполняется максимум maxConcurrent проверок.
*/
