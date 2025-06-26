package Factory5

import "fmt"

func main() {
	nextVal := getCounter()
	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())
}

func getCounter() func() int {
	// было
	//count := 0
	//
	//return func() int {
	//	count++
	//	return count
	//}

	// решение
	ch := make(chan int)
	go func() {
		count := 0
		for {
			count++
			ch <- count
		}
	}()

	return func() int {
		return <-ch
	}
}
