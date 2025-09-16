package burgerKing

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	{
		rpcCall := func() int {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			return 42
		}
		job := func() {
			result := rpcCall()
			fmt.Println(result)
		}
		job()
	}

	fmt.Println("_____")
}

/*
🔍 Что делает:
Внутри блока создаётся функция rpcCall, которая "засыпает" на 0, 1 или 2 секунды (случайно).
Затем создаётся job, которая вызывает rpcCall и печатает результат (42).
job() вызывается.
После этого программа печатает "_____".
*/
