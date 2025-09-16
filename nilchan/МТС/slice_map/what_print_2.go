package МТС

import (
	"fmt"
)

// Что выведет код?
func main() {
	var foo []int
	var bar []int
	foo = append(foo, 1)  // 1 cap=1
	foo = append(foo, 2)  // 12 cap=2
	foo = append(foo, 3)  // 123 cap=4
	bar = append(foo, 4)  // 1234 cap=4
	foo = append(foo, 5)  // 1235 cap=4 изменит тот же массив
	fmt.Println(foo, bar) // foo-1235 bar-1235
}
