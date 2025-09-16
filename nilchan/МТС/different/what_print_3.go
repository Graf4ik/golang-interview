package МТС

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 40}
	b := make([]*int, len(a))
	for i, v := range a {
		b[i] = &v
	}
	fmt.Println(*b[0], *b[1]) // 10 20
}
