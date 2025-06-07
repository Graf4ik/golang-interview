package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//===========================================================
//Задача 8
//1. Что выведется и как исправить?
//===========================================================

/*
Текущий код содержит две основные ошибки:
Нет ожидания завершения горутин.
Нет синхронизации при доступе к counter.
Либо через mutex либо через atomic
*/

func main() {
	// var counter int
	var counter int32
	wg := sync.WaitGroup{}
	// mu := sync.Mutex{}

	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		mu.Lock()
	//		counter++
	//		mu.Unlock()
	//	}()
	//}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
