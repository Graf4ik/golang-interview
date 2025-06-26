package lamoda

import (
	"fmt"
	"sync"
)

// что выведет
type listSKU []string

func (l *listSKU) getLastSKU() string {
	//return l[len(l)]
	// l — это указатель на listSKU (то есть *listSKU)
	// len(l) — это нельзя делать, потому что len() не применяется к указателям

	//чтобы сработало
	return (*l)[len(*l)-1]
}

func main() {
	items := listSKU{
		"MP990099991",
		"MP990000002",
		"MP990000003",
		"MP990000004",
		"MP990000005",
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		lastItem := items.getLastSKU()
		fmt.Printf("Last SKU is : %s\n", lastItem)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Program completed.")
} // invalid argument: len(l) (type *listSKU) is not an array, slice, string, or map
