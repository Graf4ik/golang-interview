package concurency

import (
	"context"
	"sync"
)

//===========================================================
//Задача 10
//1. Merge n channels
//2. Если один из входных каналов закрывается, то нужно закрыть все остальные каналы
//===========================================================

func case3(channels ...chan int) chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)

		for _, c := range channels {
			select {
			case _, ok := <-c:
				if ok {
					close(c)
				}
			default:
				close(c)
			}
		}
	}()

	return out
}

func mergeChannelsWithCancel(channels ...chan int) chan int {
	out := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	forward := func(c chan int) {
		defer wg.Done()
		for {
			select {
			case num, ok := <-c:
				if !ok {
					cancel() // Инициируем отмену при закрытии любого канала
					return
				}
				out <- num
			case <-ctx.Done():
				return
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go forward(c)
	}

	go func() {
		wg.Wait()
		close(out)
		cancel()
		for _, c := range channels {
			close(c)
		}
	}()

	return out
}
