package main

import "sync"

//===========================================================
//Задача 16
//===========================================================

// Написать код функции, которая делает merge N каналов. Весь входной поток перенаправляется в один канал.

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func() {
			defer wg.Done()
			for ch := range c {
				out <- ch
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
