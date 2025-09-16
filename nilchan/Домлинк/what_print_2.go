package Домлинк

import "fmt"

func prnt(a rune) {
	fmt.Println(a)
}

// что будет выведено на экран?
func main() {
	a := []rune{'a', 'b', 'c'}
	for i := range a {
		go prnt(i)
	}
	fmt.Println("Done")
}

// Done
// 0
// 2
// 1
// Порядок 0 1 2 может быть произвольным
