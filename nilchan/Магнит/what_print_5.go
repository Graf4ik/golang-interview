package Магнит

import "fmt"

// Какое поведение?
func work(done chan struct{}, out chan int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	done <- struct{}{}
}

func main() {
	out := make(chan int)
	done := make(chan struct{})
	go work(done, out)

	<-done

	for n := range out {
		fmt.Println(n)
	}
}

// ❗ Где проблема?
// Канал out — небуферизированный, а значит:
// work() не может отправить значение в канал out, пока кто-то не начнёт читать.
// Но main() не читает out, пока не получит done.
// work() не может завершить и послать done, пока не отправит все 10 чисел в out.

//👉 Получается взаимная блокировка:
// main ждёт done
// work ждёт, пока main начнёт читать из out

// ✅ Как это исправить?
// ✅ Вариант 1: Закрыть канал out, когда всё отправлено:
// func work(out chan int) {
//	 for i := 1; i <= 10; i++ {
//		 out <- i
//	 }
//	 close(out)
// }
//
// func main() {
//	 out := make(chan int)
//	 go work(out)
//
// 	for n := range out {
//	 	fmt.Println(n)
//	 }
//  }
// ✅ Вариант 2: Читать out в отдельной горутине и обрабатывать done как завершение:
