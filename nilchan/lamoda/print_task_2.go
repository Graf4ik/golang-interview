package lamoda

import (
	"fmt"
	"runtime"
	"time"
)

// что выведете программа
func main() {
	runtime.GOMAXPROCS(1)

	var n int
	go func() {
		for {
			n++
			runtime.Gosched() // отдать управление другим
		}
	}()

	time.Sleep(500 * time.Second)
	fmt.Println("Done")
}

/*
🤔 Почему программа не выведет Done?
У тебя только 1 поток (GOMAXPROCS = 1)

Горутина go func() начинает выполняться и не отдаёт управление — потому что в её бесконечном цикле нет
ни одной точки, где планировщик может переключиться

Горутина main() не может продолжить после time.Sleep, потому что её просто никогда не планируют
💡 Это называется "goroutine starvation" — одна горутина заняла весь планировщик.
*/
