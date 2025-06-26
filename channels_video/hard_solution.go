package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func processData(val int) int {
	time.Sleep(time.Duration(val) * time.Second)
	return val * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()

	now := time.Now()
	processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

// операция должна выполнится не более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-in:
					if !ok {
						return
					}
					res := processData(val)
					select {
					case out <- res:
					case <-ctx.Done():
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}
