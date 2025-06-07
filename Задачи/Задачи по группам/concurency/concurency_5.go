package main

import (
	"fmt"
	"time"
)

//===========================================================
//Задача 5
//1. Как будет работать код?
//2. Как сделать так, чтобы выводился только первый ch?
//===========================================================

func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	// Вариант 1: запускать только первую горутину (остальные не запускать)
	// Вариант 2: задержать остальные горутины с помощью time.Sleep
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- true
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- true
	}()
	go func() {
		ch3 <- true
	}()

	/*
	select выбирает случайный case, если несколько каналов готовы одновременно.
	Поэтому результат будет недетерминированным: на каждом запуске может вывестись:
	*/
	select {
	case <-ch:
		fmt.Printf("val from ch")
	case <-ch2:
		fmt.Printf("val from ch2")
	case <-ch3:
		fmt.Printf("val from ch3")
	}
}
