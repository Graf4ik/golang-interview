package Uzum

import (
	"fmt"
	"time"
)

// что выведет
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		panic(123)
	}()

	time.Sleep(time.Hour)
}

// Этот код вызовет фатальную ошибку и не выведет 123, несмотря на recover().
// Почему?
// В Go recover() может перехватывать панику только внутри той же горутины, в которой произошёл panic.

/*
Как поймать panic в другой горутине?
Чтобы безопасно обрабатывать панику в новой горутине, оберните тело функции этой горутины в defer/recover:
go func() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in goroutine:", err)
		}
	}()

	panic(123)
}()
*/
