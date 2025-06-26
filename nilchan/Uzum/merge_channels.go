package Uzum

import "sync"

func merge(ch ...chan int) chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(ch))

	for _, channel := range ch {
		go func() {
			defer wg.Done()
			for v := range channel {
				merged <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
