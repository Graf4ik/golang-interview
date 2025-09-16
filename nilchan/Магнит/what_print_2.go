package Магнит

import (
	"fmt"
	"time"
)

// Какое поведение у кода ниже?
const workers = 1000

var count int

func networkRequest() {
	time.Sleep(time.Millisecond)
}

func main() {
	for i := 0; i < workers; i++ {
		go networkRequest()
	}
	fmt.Println(count)
} // 0

// Почему?
// Функция networkRequest() делает только time.Sleep(...), не изменяя переменную count.
// Переменная count нигде не инкрементируется, не устанавливается — вообще не используется.
// Следовательно, count всегда остаётся равным 0.
