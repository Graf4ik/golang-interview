package Домлинк

import "fmt"

func foo() string {
	fmt.Println("foo")
	return "close"
}

func close(s string) {
	fmt.Println(s)
}

// что будет выведено на экран?
func main() {
	defer close(foo())
	fmt.Println("start")
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// foo start 3 2 1 0 close

// Сначала выполнится foo(), потому что аргумент функции close должен быть вычислен сразу, при попадании строки в defer.
// foo() возвращает строку "close", и происходит отложенный вызов close("close"), который выполнится в самом конце main(), после остальных defer.
