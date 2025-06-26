package main

//func processData(val int) int {
//	time.Sleep(time.Duration(val) * time.Second)
//	return val * 2
//}
//
//func main() {
//	in := make(chan int)
//	out := make(chan int)
//
//	go func() {
//		for i := range 100 {
//			in <- i
//		}
//		close(in)
//	}()
//
//	now := time.Now()
//	processParallel(in, out, 5)
//
//	for val := range out {
//		fmt.Println(val)
//	}
//	fmt.Println(time.Since(now))
//}
//
//// операция должна выполнится не более 5 секунд
//func processParallel(in <-chan int, out chan<- int, numWorkers int) {
//
//}
