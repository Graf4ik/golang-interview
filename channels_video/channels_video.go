package main

import (
	"fmt"
	"time"
)

// Написать 3 функции:
// writer - генерит числа от 1 до 10
// doubler - умножает числа на 2, имитируя работу (500мс)
// reader - читает и выводит на экран
func main() {
	reader(double(writer()))
}

func writer() <-chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for val := range in {
			time.Sleep(500 * time.Millisecond)
			out <- val * 2
		}
		close(out)
	}()

	return out
}

func reader(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}
