package main

import "fmt"

// Calling g.
// Printing in g 0
// Printing in g 1
// Printing in g 2
// Printing in g 3
// Panicking!
// Defer in g 3
// Defer in g 2
// Defer in g 1
// Recovered in f (0x137000,0xc000116060)
// Returned normally from f

// Что выведет
func main() {
	f()
	println("Returned normally from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered in f", r)
		}
	}()

	println("Calling g.")
	g(0)
	println("Returned normally from g")
}

func g(i int) {
	if i > 3 {
		println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer println("Defer in g", i)
	println("Printing in g", i)
	g(i + 1)
}

// 🔍 Что происходит по шагам
// main() вызывает f().
// f():
// Устанавливает defer с recover().
// Печатает Calling g.
// Вызывает g(0)
// g(i) — рекурсивно вызывает саму себя, пока i > 3:
// i = 0: печатает "Printing in g 0", устанавливает defer, вызывает g(1)
// i = 1: то же самое
// i = 2: то же самое
// i = 3: то же самое
// i = 4: печатает "Panicking!", вызывает panic(fmt.Sprintf("%v", i))
//
// 📌 Что делает panic
// Когда вызывается panic(...), программа начинает выход из всех функций вверх по стеку вызовов.
// При этом выполняются все defer, если они есть.
// Если находит recover(), то panic подавляется.
//
// 🧠 Важно
// Все defer println(...) внутри g(...) будут вызваны в обратном порядке (от последнего вызова к первому).
// recover() сработает в f(), остановит панику и выведет сообщение.
