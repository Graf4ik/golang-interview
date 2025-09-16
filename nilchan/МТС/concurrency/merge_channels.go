package МТС

import "sync"

// Написать код функции, которая делает merge N каналов. Весь входной
// поток перенаправляется в один канал.
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(cs))

	for _, c := range cs {
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
