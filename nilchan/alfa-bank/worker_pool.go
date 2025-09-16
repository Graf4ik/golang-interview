package alfaBank

// реализовать worker pool
// Есть 10 задач (функций), каждая засыпает на 1 сек и выводит номер воркера, который эту задачу исполнил.
// Кол-во воркеров задается при запуске.
//func main() {
//	var wg sync.WaitGroup
//	n := 3
//	c := make(chan struct{}, n)
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//
//		go func() {
//			c <- struct{}{}
//			defer wg.Done()
//			defer func() { <-c }()
//			time.Sleep(1 * time.Second)
//		}()
//	}
//
//	wg.Wait()
//}
