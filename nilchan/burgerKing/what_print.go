package burgerKing

import (
	"fmt"
	"sync"
)

func main() {
	{
		n := 10
		wg := &sync.WaitGroup{}
		wg.Add(n)
		for i := range n {
			i := i // захватываем переменную i отдельно для каждой горутины
			go func() {
				fmt.Println(i * 10) // Порядок рандомный
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
