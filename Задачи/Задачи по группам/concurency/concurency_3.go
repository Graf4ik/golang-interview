package concurency

import "fmt"

// ===========================================================
// Задача 3
// Будет ошибка что все горутины заблокированы. Какие горутины будут заблокированы? И почему?
// ===========================================================

/*
 Важное замечание:
ch <- 1 — блокирует main, потому что канал небуферизированный и никто ещё не слушает его.
А go func(...) — даже не начнёт выполняться, потому что main уже заблокирован и не дойдёт до создания горутины.
В итоге единственная активная горутина (main) уже стоит на блокирующей операции, а других ещё не существует.
*/

func main() {
	ch := make(chan int)
	// ch <- 1 // ➋ main блокируется здесь, ожидая, что кто-то примет значение из канала
	go func() { // ➌ эта горутина НИКОГДА НЕ ЗАПУСТИТСЯ
		fmt.Println(<-ch)
	}()

	ch <- 1
}
