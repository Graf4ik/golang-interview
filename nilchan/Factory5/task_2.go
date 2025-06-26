// package Factory5
package Factory5

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var wg sync.WaitGroup

	for _, val := range values {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v)
		}(val)
	}
	wg.Wait()
	time.Sleep(5 * time.Second)
}
