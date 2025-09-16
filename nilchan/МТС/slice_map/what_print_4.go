package МТС

import (
	"fmt"
)

// Что выведет код?
func main() {
	var m map[string]int // == nil
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
} // panic: assignment to entry in nil map

// 🛠 Как исправить:
// Нужно инициализировать мапу перед использованием:
// m := make(map[string]int)
