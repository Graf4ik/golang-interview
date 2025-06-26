package Uzum

import "fmt"

// что выведет
func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	a := m["one"]
	fmt.Println(&a)
}
