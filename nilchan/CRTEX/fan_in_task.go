package CRTEX

import (
	"fmt"
	"sync"
)

// Необходимо написать программу, в которой N горутин одновременно пишут в канал любое число
// а main-горутина находит сумму всех чисел, записанных в канал

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ch <- v
		}(i)
	}

	go func() {
		close(ch)
		wg.Wait()
	}()

	var sum int
	for v := range ch {
		sum += v
	}

	fmt.Printf("result %d\n", sum)
}
