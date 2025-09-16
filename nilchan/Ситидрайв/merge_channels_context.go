package Ситидрайв

import (
	"context"
	"sync"
)

// 1. Объединить все каналы в 1
// 2. Если 1 закрывается, все остальные закрыть

func merge(channels ...chan int) <-chan int {
	out := make(chan int, len(channels))
	wg := sync.WaitGroup{}
	once := sync.Once{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(len(channels))

	for _, channel := range channels {
		go func(channel chan int) {
			defer wg.Done()

			for {
				select {
				case val, ok := <-channel:
					if !ok {
						once.Do(cancel)
						return
					}

					select {
					case out <- val:
					case <-ctx.Done():
						close(channel)
						return
					}
				case <-ctx.Done():
					close(channel)
					return
				}
			}
		}(channel)
	}

	go func() {
		close(out)
		wg.Wait()
	}()

	return out
}
