package smartWay

import "fmt"

// что выведет в консоль
func main() {
	intKeys := map[int]string{1: "a", 2: "b", 3: "c"}
	_ = &intKeys[1] // что выведет
	// invalid operation: cannot take address of intKeys[1] (map index expression of type string)
}
